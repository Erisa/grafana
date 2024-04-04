//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1.DataSourceApiServer":          schema_pkg_apis_query_v0alpha1_DataSourceApiServer(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1.DataSourceApiServerList":      schema_pkg_apis_query_v0alpha1_DataSourceApiServerList(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1.QueryDataRequest":             schema_pkg_apis_query_v0alpha1_QueryDataRequest(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1.QueryDataResponse":            schema_pkg_apis_query_v0alpha1_QueryDataResponse(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1.QueryTypeDefinition":          schema_pkg_apis_query_v0alpha1_QueryTypeDefinition(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1.QueryTypeDefinitionList":      schema_pkg_apis_query_v0alpha1_QueryTypeDefinitionList(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Position":            schema_apis_query_v0alpha1_template_Position(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.QueryTemplate":       schema_apis_query_v0alpha1_template_QueryTemplate(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Target":              schema_apis_query_v0alpha1_template_Target(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.TemplateVariable":    schema_apis_query_v0alpha1_template_TemplateVariable(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.VariableReplacement": schema_apis_query_v0alpha1_template_VariableReplacement(ref),
		"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.replacement":         schema_apis_query_v0alpha1_template_replacement(ref),
	}
}

func schema_pkg_apis_query_v0alpha1_DataSourceApiServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "The data source resource is a reflection of the individual datasource instances that are exposed in the groups: {datasource}.datasource.grafana.app The status is updated periodically. The name is the plugin id",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"title": {
						SchemaProps: spec.SchemaProps{
							Description: "The display name",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Describe the plugin",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"groupVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "The group + preferred version",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"aliasIDs": {
						SchemaProps: spec.SchemaProps{
							Description: "Possible alternative plugin IDs",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"title", "groupVersion"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_query_v0alpha1_DataSourceApiServerList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "List of datasource plugins",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1.DataSourceApiServer"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/query/v0alpha1.DataSourceApiServer", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_query_v0alpha1_QueryDataRequest(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Generic query request with shared time across all values Copied from: https://github.com/grafana/grafana/blob/main/pkg/api/dtos/models.go#L62",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"from": {
						SchemaProps: spec.SchemaProps{
							Description: "From is the start time of the query.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"to": {
						SchemaProps: spec.SchemaProps{
							Description: "To is the end time of the query.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"queries": {
						SchemaProps: spec.SchemaProps{
							Description: "Datasource queries",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery"),
									},
								},
							},
						},
					},
					"debug": {
						SchemaProps: spec.SchemaProps{
							Description: "Optionally include debug information in the response",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"from", "to", "queries"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery"},
	}
}

func schema_pkg_apis_query_v0alpha1_QueryDataResponse(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Wraps backend.QueryDataResponse, however it includes TypeMeta and implements runtime.Object",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"results": {
						SchemaProps: spec.SchemaProps{
							Description: "Responses is a map of RefIDs (Unique Query ID) to *DataResponse.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana-plugin-sdk-go/backend.DataResponse"),
									},
								},
							},
						},
					},
				},
				Required: []string{"results"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana-plugin-sdk-go/backend.DataResponse"},
	}
}

func schema_pkg_apis_query_v0alpha1_QueryTypeDefinition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Defines a query behavior in a datasource.  This is a similar model to a CRD where the payload describes a valid query",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.QueryTypeDefinitionSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.QueryTypeDefinitionSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_query_v0alpha1_QueryTypeDefinitionList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1.QueryTypeDefinition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/query/v0alpha1.QueryTypeDefinition", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_apis_query_v0alpha1_template_Position(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Position is where to do replacement in the targets during render.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"start": {
						SchemaProps: spec.SchemaProps{
							Description: "Start is the byte offset within TargetKey's property of the variable. It is the start location for replacements).",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"end": {
						SchemaProps: spec.SchemaProps{
							Description: "End is the byte offset of the end of the variable.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
				},
				Required: []string{"start", "end"},
			},
		},
	}
}

func schema_apis_query_v0alpha1_template_QueryTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"title": {
						SchemaProps: spec.SchemaProps{
							Description: "A display name",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Longer description for why it is interesting",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"vars": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"key",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The variables that can be used to render",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.TemplateVariable"),
									},
								},
							},
						},
					},
					"targets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Output variables",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Target"),
									},
								},
							},
						},
					},
				},
				Required: []string{"targets"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Target", "github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.TemplateVariable"},
	}
}

func schema_apis_query_v0alpha1_template_Target(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"dataType": {
						SchemaProps: spec.SchemaProps{
							Description: "DataType is the returned Dataplane type from the query.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"variables": {
						SchemaProps: spec.SchemaProps{
							Description: "Variables that will be replaced in the query",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Default: map[string]interface{}{},
													Ref:     ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.VariableReplacement"),
												},
											},
										},
									},
								},
							},
						},
					},
					"properties": {
						SchemaProps: spec.SchemaProps{
							Description: "Query target",
							Ref:         ref("github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery"),
						},
					},
				},
				Required: []string{"variables", "properties"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana-plugin-sdk-go/experimental/apis/data/v0alpha1.DataQuery", "github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.VariableReplacement"},
	}
}

func schema_apis_query_v0alpha1_template_TemplateVariable(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TemplateVariable is the definition of a variable that will be interpolated in targets.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"key": {
						SchemaProps: spec.SchemaProps{
							Description: "Key is the name of the variable.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"defaultValues": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "DefaultValue is the value to be used when there is no selected value during render.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"valueListDefinition": {
						SchemaProps: spec.SchemaProps{
							Description: "ValueListDefinition is the object definition used by the FE to get a list of possible values to select for render.",
							Ref:         ref("github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
				Required: []string{"key"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apimachinery/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_apis_query_v0alpha1_template_VariableReplacement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QueryVariable is the definition of a variable that will be interpolated in targets.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path is the location of the property within a target. The format for this is not figured out yet (Maybe JSONPath?). Idea: [\"string\", int, \"string\"] where int indicates array offset",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"position": {
						SchemaProps: spec.SchemaProps{
							Description: "Positions is a list of where to perform the interpolation within targets during render. The first string is the Idx of the target as a string, since openAPI does not support ints as map keys",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Position"),
						},
					},
					"format": {
						SchemaProps: spec.SchemaProps{
							Description: "How values should be interpolated\n\nPossible enum values:\n - `\"csv\"` Formats variables with multiple values as a comma-separated string.\n - `\"doublequote\"` Formats single- and multi-valued variables into a comma-separated string\n - `\"json\"` Formats variables with multiple values as a comma-separated string.\n - `\"pipe\"` Formats variables with multiple values into a pipe-separated string.\n - `\"raw\"` Formats variables with multiple values into comma-separated string. This is the default behavior when no format is specified\n - `\"singlequote\"` Formats single- and multi-valued variables into a comma-separated string",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"csv", "doublequote", "json", "pipe", "raw", "singlequote"},
						},
					},
				},
				Required: []string{"path"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Position"},
	}
}

func schema_apis_query_v0alpha1_template_replacement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"Position": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Position"),
						},
					},
					"TemplateVariable": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.TemplateVariable"),
						},
					},
					"format": {
						SchemaProps: spec.SchemaProps{
							Description: "Possible enum values:\n - `\"csv\"` Formats variables with multiple values as a comma-separated string.\n - `\"doublequote\"` Formats single- and multi-valued variables into a comma-separated string\n - `\"json\"` Formats variables with multiple values as a comma-separated string.\n - `\"pipe\"` Formats variables with multiple values into a pipe-separated string.\n - `\"raw\"` Formats variables with multiple values into comma-separated string. This is the default behavior when no format is specified\n - `\"singlequote\"` Formats single- and multi-valued variables into a comma-separated string",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"csv", "doublequote", "json", "pipe", "raw", "singlequote"},
						},
					},
				},
				Required: []string{"Position", "TemplateVariable", "format"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.Position", "github.com/grafana/grafana/pkg/apis/query/v0alpha1/template.TemplateVariable"},
	}
}
