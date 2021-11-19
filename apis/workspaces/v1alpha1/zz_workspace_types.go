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

type WorkspaceObservation struct {
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkspaceParameters struct {

	// +kubebuilder:validation:Required
	BundleID *string `json:"bundleId" tf:"bundle_id,omitempty"`

	// +kubebuilder:validation:Required
	DirectoryID *string `json:"directoryId" tf:"directory_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RootVolumeEncryptionEnabled *bool `json:"rootVolumeEncryptionEnabled,omitempty" tf:"root_volume_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	UserName *string `json:"userName" tf:"user_name,omitempty"`

	// +kubebuilder:validation:Optional
	UserVolumeEncryptionEnabled *bool `json:"userVolumeEncryptionEnabled,omitempty" tf:"user_volume_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeEncryptionKey *string `json:"volumeEncryptionKey,omitempty" tf:"volume_encryption_key,omitempty"`

	// +kubebuilder:validation:Optional
	WorkspaceProperties []WorkspacePropertiesParameters `json:"workspaceProperties,omitempty" tf:"workspace_properties,omitempty"`
}

type WorkspacePropertiesObservation struct {
}

type WorkspacePropertiesParameters struct {

	// +kubebuilder:validation:Optional
	ComputeTypeName *string `json:"computeTypeName,omitempty" tf:"compute_type_name,omitempty"`

	// +kubebuilder:validation:Optional
	RootVolumeSizeGib *int64 `json:"rootVolumeSizeGib,omitempty" tf:"root_volume_size_gib,omitempty"`

	// +kubebuilder:validation:Optional
	RunningMode *string `json:"runningMode,omitempty" tf:"running_mode,omitempty"`

	// +kubebuilder:validation:Optional
	RunningModeAutoStopTimeoutInMinutes *int64 `json:"runningModeAutoStopTimeoutInMinutes,omitempty" tf:"running_mode_auto_stop_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	UserVolumeSizeGib *int64 `json:"userVolumeSizeGib,omitempty" tf:"user_volume_size_gib,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workspace is the Schema for the Workspaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
