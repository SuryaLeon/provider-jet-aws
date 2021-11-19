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

type CloudwatchEncryptionObservation struct {
}

type CloudwatchEncryptionParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchEncryptionMode *string `json:"cloudwatchEncryptionMode,omitempty" tf:"cloudwatch_encryption_mode,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Required
	CloudwatchEncryption []CloudwatchEncryptionParameters `json:"cloudwatchEncryption" tf:"cloudwatch_encryption,omitempty"`

	// +kubebuilder:validation:Required
	JobBookmarksEncryption []JobBookmarksEncryptionParameters `json:"jobBookmarksEncryption" tf:"job_bookmarks_encryption,omitempty"`

	// +kubebuilder:validation:Required
	S3Encryption []S3EncryptionParameters `json:"s3Encryption" tf:"s3_encryption,omitempty"`
}

type JobBookmarksEncryptionObservation struct {
}

type JobBookmarksEncryptionParameters struct {

	// +kubebuilder:validation:Optional
	JobBookmarksEncryptionMode *string `json:"jobBookmarksEncryptionMode,omitempty" tf:"job_bookmarks_encryption_mode,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type S3EncryptionObservation struct {
}

type S3EncryptionParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	S3EncryptionMode *string `json:"s3EncryptionMode,omitempty" tf:"s3_encryption_mode,omitempty"`
}

type SecurityConfigurationObservation struct {
}

type SecurityConfigurationParameters struct {

	// +kubebuilder:validation:Required
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration" tf:"encryption_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// SecurityConfigurationSpec defines the desired state of SecurityConfiguration
type SecurityConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityConfigurationParameters `json:"forProvider"`
}

// SecurityConfigurationStatus defines the observed state of SecurityConfiguration.
type SecurityConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfiguration is the Schema for the SecurityConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityConfigurationSpec   `json:"spec"`
	Status            SecurityConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityConfigurationList contains a list of SecurityConfigurations
type SecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityConfiguration `json:"items"`
}

// Repository type metadata.
var (
	SecurityConfiguration_Kind             = "SecurityConfiguration"
	SecurityConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityConfiguration_Kind}.String()
	SecurityConfiguration_KindAPIVersion   = SecurityConfiguration_Kind + "." + CRDGroupVersion.String()
	SecurityConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(SecurityConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityConfiguration{}, &SecurityConfigurationList{})
}
