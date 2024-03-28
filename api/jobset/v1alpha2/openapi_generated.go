//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.FailurePolicy":       schema_jobset_api_jobset_v1alpha2_FailurePolicy(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSet":              schema_jobset_api_jobset_v1alpha2_JobSet(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetList":          schema_jobset_api_jobset_v1alpha2_JobSetList(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetSpec":          schema_jobset_api_jobset_v1alpha2_JobSetSpec(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetStatus":        schema_jobset_api_jobset_v1alpha2_JobSetStatus(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.Network":             schema_jobset_api_jobset_v1alpha2_Network(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.ReplicatedJob":       schema_jobset_api_jobset_v1alpha2_ReplicatedJob(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.ReplicatedJobStatus": schema_jobset_api_jobset_v1alpha2_ReplicatedJobStatus(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.StartupPolicy":       schema_jobset_api_jobset_v1alpha2_StartupPolicy(ref),
		"sigs.k8s.io/jobset/api/jobset/v1alpha2.SuccessPolicy":       schema_jobset_api_jobset_v1alpha2_SuccessPolicy(ref),
	}
}

func schema_jobset_api_jobset_v1alpha2_FailurePolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"maxRestarts": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxRestarts defines the limit on the number of JobSet restarts. A restart is achieved by recreating all active child jobs.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
	}
}

func schema_jobset_api_jobset_v1alpha2_JobSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobSet is the Schema for the jobsets API",
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
							Ref:     ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetSpec", "sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSetStatus"},
	}
}

func schema_jobset_api_jobset_v1alpha2_JobSetList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobSetList contains a list of JobSet",
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
										Ref:     ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSet"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "sigs.k8s.io/jobset/api/jobset/v1alpha2.JobSet"},
	}
}

func schema_jobset_api_jobset_v1alpha2_JobSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobSetSpec defines the desired state of JobSet",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicatedJobs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ReplicatedJobs is the group of jobs that will form the set.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.ReplicatedJob"),
									},
								},
							},
						},
					},
					"network": {
						SchemaProps: spec.SchemaProps{
							Description: "Network defines the networking options for the jobset.",
							Ref:         ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.Network"),
						},
					},
					"successPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "SuccessPolicy configures when to declare the JobSet as succeeded. The JobSet is always declared succeeded if all jobs in the set finished with status complete.",
							Ref:         ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.SuccessPolicy"),
						},
					},
					"failurePolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "FailurePolicy, if set, configures when to declare the JobSet as failed. The JobSet is always declared failed if any job in the set finished with status failed.",
							Ref:         ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.FailurePolicy"),
						},
					},
					"startupPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "StartupPolicy, if set, configures in what order jobs must be started",
							Ref:         ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.StartupPolicy"),
						},
					},
					"suspend": {
						SchemaProps: spec.SchemaProps{
							Description: "Suspend suspends all running child Jobs when set to true.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"managedBy": {
						SchemaProps: spec.SchemaProps{
							Description: "ManagedBy is used to indicate the controller or entity that manages a JobSet",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"sigs.k8s.io/jobset/api/jobset/v1alpha2.FailurePolicy", "sigs.k8s.io/jobset/api/jobset/v1alpha2.Network", "sigs.k8s.io/jobset/api/jobset/v1alpha2.ReplicatedJob", "sigs.k8s.io/jobset/api/jobset/v1alpha2.StartupPolicy", "sigs.k8s.io/jobset/api/jobset/v1alpha2.SuccessPolicy"},
	}
}

func schema_jobset_api_jobset_v1alpha2_JobSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobSetStatus defines the observed state of JobSet",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
					"restarts": {
						SchemaProps: spec.SchemaProps{
							Description: "Restarts tracks the number of times the JobSet has restarted (i.e. recreated in case of RecreateAll policy).",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"replicatedJobsStatus": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ReplicatedJobsStatus track the number of JobsReady for each replicatedJob.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("sigs.k8s.io/jobset/api/jobset/v1alpha2.ReplicatedJobStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Condition", "sigs.k8s.io/jobset/api/jobset/v1alpha2.ReplicatedJobStatus"},
	}
}

func schema_jobset_api_jobset_v1alpha2_Network(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enableDNSHostnames": {
						SchemaProps: spec.SchemaProps{
							Description: "EnableDNSHostnames allows pods to be reached via their hostnames. Pods will be reachable using the fully qualified pod hostname: <jobSet.name>-<spec.replicatedJob.name>-<job-index>-<pod-index>.<subdomain>",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"subdomain": {
						SchemaProps: spec.SchemaProps{
							Description: "Subdomain is an explicit choice for a network subdomain name When set, any replicated job in the set is added to this network. Defaults to <jobSet.name> if not set.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_jobset_api_jobset_v1alpha2_ReplicatedJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the entry and will be used as a suffix for the Job name.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template defines the template of the Job that will be created.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/api/batch/v1.JobTemplateSpec"),
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the number of jobs that will be created from this ReplicatedJob's template. Jobs names will be in the format: <jobSet.name>-<spec.replicatedJob.name>-<job-index>",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"name", "template"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/batch/v1.JobTemplateSpec"},
	}
}

func schema_jobset_api_jobset_v1alpha2_ReplicatedJobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReplicatedJobStatus defines the observed ReplicatedJobs Readiness.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the ReplicatedJob.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "Ready is the number of child Jobs where the number of ready pods and completed pods is greater than or equal to the total expected pod count for the Job (i.e., the minimum of job.spec.parallelism and job.spec.completions).",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"succeeded": {
						SchemaProps: spec.SchemaProps{
							Description: "Succeeded is the number of successfully completed child Jobs.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"failed": {
						SchemaProps: spec.SchemaProps{
							Description: "Failed is the number of failed child Jobs.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"active": {
						SchemaProps: spec.SchemaProps{
							Description: "Active is the number of child Jobs with at least 1 pod in a running or pending state which are not marked for deletion.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"suspended": {
						SchemaProps: spec.SchemaProps{
							Description: "Suspended is the number of child Jobs which are in a suspended state.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"name", "ready", "succeeded", "failed", "active", "suspended"},
			},
		},
	}
}

func schema_jobset_api_jobset_v1alpha2_StartupPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"startupPolicyOrder": {
						SchemaProps: spec.SchemaProps{
							Description: "StartupPolicyOrder determines the startup order of the ReplicatedJobs. AnyOrder means to start replicated jobs in any order. InOrder means to start them as they are listed in the JobSet. A ReplicatedJob is started only when all the jobs of the previous one are ready.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"startupPolicyOrder"},
			},
		},
	}
}

func schema_jobset_api_jobset_v1alpha2_SuccessPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"operator": {
						SchemaProps: spec.SchemaProps{
							Description: "Operator determines either All or Any of the selected jobs should succeed to consider the JobSet successful",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"targetReplicatedJobs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "TargetReplicatedJobs are the names of the replicated jobs the operator will apply to. A null or empty list will apply to all replicatedJobs.",
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
				Required: []string{"operator"},
			},
		},
	}
}
