package remote

import (
	"context"
	"fmt"

	"github.com/grafana/grafana/pkg/infra/log"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/services/ngalert/notifier"
)

type Mode int

// NOTE: we don't really need ModeRemoteOnly in the forked Alertmanager.
const (
	ModeRemoteSecondary Mode = iota
	ModeRemotePrimary
)

type forkedAlertmanager struct {
	log        log.Logger
	internal   notifier.Alertmanager
	remote     *Alertmanager
	configSent bool
	mode       Mode
}

func NewForkedAlertmanager(internal notifier.Alertmanager, remote *Alertmanager, m Mode) *forkedAlertmanager {
	return &forkedAlertmanager{
		internal: internal,
		remote:   remote,
		mode:     m,
		log:      log.New("ngalert.remote.forked-alertmanager"),
	}
}

// Note: this is called on startup and on sync.
func (fam *forkedAlertmanager) ApplyConfig(ctx context.Context, config *models.AlertConfiguration) error {
	if fam.mode == ModeRemoteSecondary && !fam.configSent {
		// If we're in Internal Primary mode we just send the config on startup.
		if err := fam.remote.ApplyConfig(ctx, config); err != nil {
			return err
		}
		fam.configSent = true
	} else {
		// TODO: delet
		fmt.Println("tried to send config but config was already sent")
	}
	return fam.internal.ApplyConfig(ctx, config)
}

func (fam *forkedAlertmanager) SaveAndApplyConfig(ctx context.Context, config *apimodels.PostableUserConfig) error {
	// NOTE: this is used in:
	//	- ActivateHistoricalConfiguration
	//	- ApplyAlertmanagerConfiguration, used in RoutePostAlertingConfig
	if fam.mode != ModeRemoteSecondary {
		if err := fam.remote.SaveAndApplyConfig(ctx, config); err != nil {
			return err
		}
	}
	return fam.internal.SaveAndApplyConfig(ctx, config)
}

func (fam *forkedAlertmanager) SaveAndApplyDefaultConfig(ctx context.Context) error {
	// NOTE: this is used in
	//	- RouteDeleteAlertingConfig
	//	- SyncAlertmanagersForOrgs when no db config is found
	if fam.mode != ModeRemoteSecondary {
		if err := fam.remote.SaveAndApplyDefaultConfig(ctx); err != nil {
			return err
		}
	}
	return fam.internal.SaveAndApplyDefaultConfig(ctx)
}

func (fam *forkedAlertmanager) GetStatus() apimodels.GettableStatus {
	if fam.mode == ModeRemoteSecondary {
		return fam.internal.GetStatus()
	}
	return fam.remote.GetStatus()
}

func (fam *forkedAlertmanager) CreateSilence(ctx context.Context, silence *apimodels.PostableSilence) (string, error) {
	id, err := fam.internal.CreateSilence(ctx, silence)
	if err != nil {
		return "", err
	}
	// In ModeRemoteSecondary we just create the silence in the internal Alertmanager.
	if fam.mode == ModeRemoteSecondary {
		return id, nil
	}

	// If were not in ModeRemoteSecodary we care about the id returned from the remote Alertmanager.
	return fam.remote.CreateSilence(ctx, silence)
}

func (fam *forkedAlertmanager) DeleteSilence(ctx context.Context, id string) error {
	if fam.mode == ModeRemotePrimary {
		if err := fam.remote.DeleteSilence(ctx, id); err != nil {
			return err
		}
	}

	return fam.internal.DeleteSilence(ctx, id)
}

func (fam *forkedAlertmanager) GetSilence(ctx context.Context, id string) (apimodels.GettableSilence, error) {
	if fam.mode != ModeRemoteSecondary {
		return fam.remote.GetSilence(ctx, id)
	}
	return fam.internal.GetSilence(ctx, id)
}

func (fam *forkedAlertmanager) ListSilences(ctx context.Context, filter []string) (apimodels.GettableSilences, error) {
	if fam.mode != ModeRemoteSecondary {
		return fam.remote.ListSilences(ctx, filter)
	}
	return fam.internal.ListSilences(ctx, filter)
}

// Alerts
func (fam *forkedAlertmanager) GetAlerts(ctx context.Context, active, silenced, inhibited bool, filter []string, receiver string) (apimodels.GettableAlerts, error) {
	if fam.mode != ModeRemoteSecondary {
		return fam.remote.GetAlerts(ctx, active, silenced, inhibited, filter, receiver)
	}
	return fam.internal.GetAlerts(ctx, active, silenced, inhibited, filter, receiver)
}

func (fam *forkedAlertmanager) GetAlertGroups(ctx context.Context, active, silenced, inhibited bool, filter []string, receiver string) (apimodels.AlertGroups, error) {
	if fam.mode != ModeRemoteSecondary {
		return fam.remote.GetAlertGroups(ctx, active, silenced, inhibited, filter, receiver)
	}
	return fam.internal.GetAlertGroups(ctx, active, silenced, inhibited, filter, receiver)
}

func (fam *forkedAlertmanager) PutAlerts(ctx context.Context, alerts apimodels.PostableAlerts) error {
	if fam.mode != ModeRemoteSecondary {
		return fam.remote.PutAlerts(ctx, alerts)
	}
	return fam.internal.PutAlerts(ctx, alerts)
}

// Receivers
func (fam *forkedAlertmanager) GetReceivers(ctx context.Context) ([]apimodels.Receiver, error) {
	if fam.mode != ModeRemoteSecondary {
		return fam.remote.GetReceivers(ctx)
	}
	return fam.internal.GetReceivers(ctx)
}

func (fam *forkedAlertmanager) TestReceivers(ctx context.Context, c apimodels.TestReceiversConfigBodyParams) (*notifier.TestReceiversResult, error) {
	// TODO: not implmented in Cloud AM
	return fam.internal.TestReceivers(ctx, c)
}

func (fam *forkedAlertmanager) TestTemplate(ctx context.Context, c apimodels.TestTemplatesConfigBodyParams) (*notifier.TestTemplatesResults, error) {
	// TODO: not implmented in Cloud AM
	return fam.internal.TestTemplate(ctx, c)
}

// State
func (fam *forkedAlertmanager) CleanUp() {
	// No cleanup to do in the remote Alertmanager.
	fam.internal.CleanUp()
}

func (fam *forkedAlertmanager) StopAndWait() {
	fam.internal.StopAndWait()
	// Stop senders.
	fam.remote.StopAndWait()
}

func (fam *forkedAlertmanager) Ready() bool {
	if ready := fam.internal.Ready(); !ready {
		return false
	}
	return fam.remote.Ready()
}
