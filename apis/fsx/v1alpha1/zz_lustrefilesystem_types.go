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

type LustreFileSystemObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	MountName *string `json:"mountName,omitempty" tf:"mount_name,omitempty"`

	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type LustreFileSystemParameters struct {

	// +kubebuilder:validation:Optional
	AutoImportPolicy *string `json:"autoImportPolicy,omitempty" tf:"auto_import_policy,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticBackupRetentionDays *int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// +kubebuilder:validation:Optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// +kubebuilder:validation:Optional
	DataCompressionType *string `json:"dataCompressionType,omitempty" tf:"data_compression_type,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// +kubebuilder:validation:Optional
	DriveCacheType *string `json:"driveCacheType,omitempty" tf:"drive_cache_type,omitempty"`

	// +kubebuilder:validation:Optional
	ExportPath *string `json:"exportPath,omitempty" tf:"export_path,omitempty"`

	// +kubebuilder:validation:Optional
	ImportPath *string `json:"importPath,omitempty" tf:"import_path,omitempty"`

	// +kubebuilder:validation:Optional
	ImportedFileChunkSize *int64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	PerUnitStorageThroughput *int64 `json:"perUnitStorageThroughput,omitempty" tf:"per_unit_storage_throughput,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	StorageCapacity *int64 `json:"storageCapacity" tf:"storage_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

// LustreFileSystemSpec defines the desired state of LustreFileSystem
type LustreFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LustreFileSystemParameters `json:"forProvider"`
}

// LustreFileSystemStatus defines the observed state of LustreFileSystem.
type LustreFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LustreFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LustreFileSystem is the Schema for the LustreFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LustreFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LustreFileSystemSpec   `json:"spec"`
	Status            LustreFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LustreFileSystemList contains a list of LustreFileSystems
type LustreFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LustreFileSystem `json:"items"`
}

// Repository type metadata.
var (
	LustreFileSystem_Kind             = "LustreFileSystem"
	LustreFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LustreFileSystem_Kind}.String()
	LustreFileSystem_KindAPIVersion   = LustreFileSystem_Kind + "." + CRDGroupVersion.String()
	LustreFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(LustreFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&LustreFileSystem{}, &LustreFileSystemList{})
}
