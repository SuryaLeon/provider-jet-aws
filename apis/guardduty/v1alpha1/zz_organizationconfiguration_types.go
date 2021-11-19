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

type DatasourcesS3LogsObservation struct {
}

type DatasourcesS3LogsParameters struct {

	// +kubebuilder:validation:Required
	AutoEnable *bool `json:"autoEnable" tf:"auto_enable,omitempty"`
}

type OrganizationConfigurationDatasourcesObservation struct {
}

type OrganizationConfigurationDatasourcesParameters struct {

	// +kubebuilder:validation:Optional
	S3Logs []DatasourcesS3LogsParameters `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type OrganizationConfigurationObservation struct {
}

type OrganizationConfigurationParameters struct {

	// +kubebuilder:validation:Required
	AutoEnable *bool `json:"autoEnable" tf:"auto_enable,omitempty"`

	// +kubebuilder:validation:Optional
	Datasources []OrganizationConfigurationDatasourcesParameters `json:"datasources,omitempty" tf:"datasources,omitempty"`

	// +kubebuilder:validation:Required
	DetectorID *string `json:"detectorId" tf:"detector_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// OrganizationConfigurationSpec defines the desired state of OrganizationConfiguration
type OrganizationConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationConfigurationParameters `json:"forProvider"`
}

// OrganizationConfigurationStatus defines the observed state of OrganizationConfiguration.
type OrganizationConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationConfiguration is the Schema for the OrganizationConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type OrganizationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationConfigurationSpec   `json:"spec"`
	Status            OrganizationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationConfigurationList contains a list of OrganizationConfigurations
type OrganizationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationConfiguration `json:"items"`
}

// Repository type metadata.
var (
	OrganizationConfiguration_Kind             = "OrganizationConfiguration"
	OrganizationConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationConfiguration_Kind}.String()
	OrganizationConfiguration_KindAPIVersion   = OrganizationConfiguration_Kind + "." + CRDGroupVersion.String()
	OrganizationConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationConfiguration{}, &OrganizationConfigurationList{})
}
