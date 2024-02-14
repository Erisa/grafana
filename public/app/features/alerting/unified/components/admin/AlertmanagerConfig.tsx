import { css } from '@emotion/css';
import React, { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import AutoSizer from 'react-virtualized-auto-sizer';

import { GrafanaTheme2 } from '@grafana/data';
import { Alert, Button, CodeEditor, Stack, useStyles2 } from '@grafana/ui';

import { useAlertmanagerConfig } from '../../hooks/useAlertmanagerConfig';
import { useUnifiedAlertingSelector } from '../../hooks/useUnifiedAlertingSelector';
import { isVanillaPrometheusAlertManagerDataSource } from '../../utils/datasource';
import { Spacer } from '../Spacer';

export interface FormValues {
  configJSON: string;
}

interface Props {
  alertmanagerName: string;
  onDismiss: () => void;
  onSave: (dataSourceName: string, oldConfig: string, newConfig: string) => void;
  onReset: (dataSourceName: string) => void;
}

export default function AlertmanagerConfig({ alertmanagerName, onDismiss, onSave, onReset }: Props): JSX.Element {
  const { loading: isDeleting, error: deletingError } = useUnifiedAlertingSelector((state) => state.deleteAMConfig);
  const { loading: isSaving, error: savingError } = useUnifiedAlertingSelector((state) => state.saveAMConfig);

  const readOnly = alertmanagerName ? isVanillaPrometheusAlertManagerDataSource(alertmanagerName) : false;
  const styles = useStyles2(getStyles);

  const {
    currentData: config,
    error: loadingError,
    isLoading: isLoadingConfig,
  } = useAlertmanagerConfig(alertmanagerName);

  const defaultValues = {
    configJSON: config ? JSON.stringify(config, null, 2) : '',
  };

  const {
    register,
    setValue,
    setError,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({
    defaultValues,
  });

  useEffect(() => {
    if (savingError) {
      setError('configJSON', { type: 'deps', message: savingError.message });
    }
  }, [savingError, setError]);

  useEffect(() => {
    if (deletingError) {
      setError('configJSON', { type: 'deps', message: deletingError.message });
    }
  }, [deletingError, setError]);

  // manually register the config field with validation
  register('configJSON', {
    required: { value: true, message: 'Required' },
    validate: (value: string) => {
      try {
        JSON.parse(value);
        return true;
      } catch (e) {
        return e instanceof Error ? e.message : 'JSON is invalid';
      }
    },
  });

  const handleSave = handleSubmit(
    (values: FormValues) => {
      onSave(alertmanagerName, defaultValues.configJSON, values.configJSON);
    },
    (errors) => {
      console.error(errors);
    }
  );

  const isOperating = isLoadingConfig || isDeleting || isSaving;

  /* loading error, if this fails don't bother rendering the form */
  if (loadingError) {
    return (
      <Alert severity="error" title="Failed to load Alertmanager configuration">
        {loadingError.message ?? 'An unkown error occurred.'}
      </Alert>
    );
  }

  return (
    <div className={styles.container}>
      {/* form error state */}
      {errors.configJSON && (
        <Alert severity="error" title="Oops, something went wrong">
          {errors.configJSON.message || 'An unknown error occurred.'}
        </Alert>
      )}

      {/* resetting state */}
      {isDeleting && (
        <Alert severity="info" title="Resetting Alertmanager configuration">
          It might take a while...
        </Alert>
      )}

      <div className={styles.content}>
        <AutoSizer>
          {({ height, width }) => (
            <>
              <CodeEditor
                language="json"
                width={width}
                height={height}
                showLineNumbers={true}
                monacoOptions={{
                  scrollBeyondLastLine: false,
                }}
                value={defaultValues.configJSON}
                showMiniMap={false}
                onSave={(value) => setValue('configJSON', value)}
                onBlur={(value) => setValue('configJSON', value)}
                readOnly={isOperating}
              />
            </>
          )}
        </AutoSizer>
      </div>
      {!readOnly && (
        <Stack justifyContent="flex-end">
          <Button variant="destructive" onClick={() => undefined} disabled={isOperating}>
            Reset
          </Button>
          <Spacer />
          <Button variant="secondary" onClick={() => onDismiss()} disabled={isOperating}>
            Cancel
          </Button>
          <Button variant="primary" onClick={handleSave} disabled={isOperating}>
            Save
          </Button>
        </Stack>
      )}
    </div>
  );
}

const getStyles = (theme: GrafanaTheme2) => ({
  container: css({
    display: 'flex',
    flexDirection: 'column',
    height: '100%',
    gap: theme.spacing(2),
  }),
  content: css({
    flex: '1 1 100%',
  }),
});
