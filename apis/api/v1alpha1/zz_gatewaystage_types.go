/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessLogSettingsObservation struct {
}

type AccessLogSettingsParameters struct {

	// +kubebuilder:validation:Required
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`

	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`
}

type GatewayStageObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ExecutionArn *string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`

	InvokeURL *string `json:"invokeUrl,omitempty" tf:"invoke_url,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GatewayStageParameters struct {

	// +kubebuilder:validation:Optional
	AccessLogSettings []AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`

	// +kubebuilder:validation:Optional
	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	CacheClusterSize *string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size,omitempty"`

	// +kubebuilder:validation:Optional
	ClientCertificateID *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id,omitempty"`

	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DocumentationVersion *string `json:"documentationVersion,omitempty" tf:"documentation_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`

	// +kubebuilder:validation:Required
	StageName *string `json:"stageName" tf:"stage_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Variables map[string]*string `json:"variables,omitempty" tf:"variables,omitempty"`

	// +kubebuilder:validation:Optional
	XrayTracingEnabled *bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled,omitempty"`
}

// GatewayStageSpec defines the desired state of GatewayStage
type GatewayStageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayStageParameters `json:"forProvider"`
}

// GatewayStageStatus defines the observed state of GatewayStage.
type GatewayStageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayStageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayStage is the Schema for the GatewayStages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayStage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayStageSpec   `json:"spec"`
	Status            GatewayStageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayStageList contains a list of GatewayStages
type GatewayStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayStage `json:"items"`
}

// Repository type metadata.
var (
	GatewayStage_Kind             = "GatewayStage"
	GatewayStage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayStage_Kind}.String()
	GatewayStage_KindAPIVersion   = GatewayStage_Kind + "." + CRDGroupVersion.String()
	GatewayStage_GroupVersionKind = CRDGroupVersion.WithKind(GatewayStage_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayStage{}, &GatewayStageList{})
}
