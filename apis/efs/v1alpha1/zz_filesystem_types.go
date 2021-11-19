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

type FileSystemObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	AvailabilityZoneID *string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	NumberOfMountTargets *int64 `json:"numberOfMountTargets,omitempty" tf:"number_of_mount_targets,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	SizeInBytes []SizeInBytesObservation `json:"sizeInBytes,omitempty" tf:"size_in_bytes,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FileSystemParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name,omitempty"`

	// +kubebuilder:validation:Optional
	CreationToken *string `json:"creationToken,omitempty" tf:"creation_token,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	LifecyclePolicy []LifecyclePolicyParameters `json:"lifecyclePolicy,omitempty" tf:"lifecycle_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PerformanceMode *string `json:"performanceMode,omitempty" tf:"performance_mode,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisionedThroughputInMibps *float64 `json:"provisionedThroughputInMibps,omitempty" tf:"provisioned_throughput_in_mibps,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ThroughputMode *string `json:"throughputMode,omitempty" tf:"throughput_mode,omitempty"`
}

type LifecyclePolicyObservation struct {
}

type LifecyclePolicyParameters struct {

	// +kubebuilder:validation:Required
	TransitionToIa *string `json:"transitionToIa" tf:"transition_to_ia,omitempty"`
}

type SizeInBytesObservation struct {
	Value *int64 `json:"value,omitempty" tf:"value,omitempty"`

	ValueInIa *int64 `json:"valueInIa,omitempty" tf:"value_in_ia,omitempty"`

	ValueInStandard *int64 `json:"valueInStandard,omitempty" tf:"value_in_standard,omitempty"`
}

type SizeInBytesParameters struct {
}

// FileSystemSpec defines the desired state of FileSystem
type FileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileSystemParameters `json:"forProvider"`
}

// FileSystemStatus defines the observed state of FileSystem.
type FileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystem is the Schema for the FileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSystemSpec   `json:"spec"`
	Status            FileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystemList contains a list of FileSystems
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystem `json:"items"`
}

// Repository type metadata.
var (
	FileSystem_Kind             = "FileSystem"
	FileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FileSystem_Kind}.String()
	FileSystem_KindAPIVersion   = FileSystem_Kind + "." + CRDGroupVersion.String()
	FileSystem_GroupVersionKind = CRDGroupVersion.WithKind(FileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&FileSystem{}, &FileSystemList{})
}
