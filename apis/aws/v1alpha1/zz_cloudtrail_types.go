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

type CloudtrailObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	HomeRegion *string `json:"homeRegion,omitempty" tf:"home_region,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CloudtrailParameters struct {

	// +kubebuilder:validation:Optional
	CloudWatchLogsGroupArn *string `json:"cloudWatchLogsGroupArn,omitempty" tf:"cloud_watch_logs_group_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn,omitempty" tf:"cloud_watch_logs_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	EnableLogFileValidation *bool `json:"enableLogFileValidation,omitempty" tf:"enable_log_file_validation,omitempty"`

	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// +kubebuilder:validation:Optional
	EventSelector []EventSelectorParameters `json:"eventSelector,omitempty" tf:"event_selector,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeGlobalServiceEvents *bool `json:"includeGlobalServiceEvents,omitempty" tf:"include_global_service_events,omitempty"`

	// +kubebuilder:validation:Optional
	InsightSelector []InsightSelectorParameters `json:"insightSelector,omitempty" tf:"insight_selector,omitempty"`

	// +kubebuilder:validation:Optional
	IsMultiRegionTrail *bool `json:"isMultiRegionTrail,omitempty" tf:"is_multi_region_trail,omitempty"`

	// +kubebuilder:validation:Optional
	IsOrganizationTrail *bool `json:"isOrganizationTrail,omitempty" tf:"is_organization_trail,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	S3BucketName *string `json:"s3BucketName" tf:"s3_bucket_name,omitempty"`

	// +kubebuilder:validation:Optional
	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	SnsTopicName *string `json:"snsTopicName,omitempty" tf:"sns_topic_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DataResourceObservation struct {
}

type DataResourceParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type EventSelectorObservation struct {
}

type EventSelectorParameters struct {

	// +kubebuilder:validation:Optional
	DataResource []DataResourceParameters `json:"dataResource,omitempty" tf:"data_resource,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeManagementEvents *bool `json:"includeManagementEvents,omitempty" tf:"include_management_events,omitempty"`

	// +kubebuilder:validation:Optional
	ReadWriteType *string `json:"readWriteType,omitempty" tf:"read_write_type,omitempty"`
}

type InsightSelectorObservation struct {
}

type InsightSelectorParameters struct {

	// +kubebuilder:validation:Required
	InsightType *string `json:"insightType" tf:"insight_type,omitempty"`
}

// CloudtrailSpec defines the desired state of Cloudtrail
type CloudtrailSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudtrailParameters `json:"forProvider"`
}

// CloudtrailStatus defines the observed state of Cloudtrail.
type CloudtrailStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudtrailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cloudtrail is the Schema for the Cloudtrails API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Cloudtrail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudtrailSpec   `json:"spec"`
	Status            CloudtrailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudtrailList contains a list of Cloudtrails
type CloudtrailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cloudtrail `json:"items"`
}

// Repository type metadata.
var (
	Cloudtrail_Kind             = "Cloudtrail"
	Cloudtrail_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cloudtrail_Kind}.String()
	Cloudtrail_KindAPIVersion   = Cloudtrail_Kind + "." + CRDGroupVersion.String()
	Cloudtrail_GroupVersionKind = CRDGroupVersion.WithKind(Cloudtrail_Kind)
)

func init() {
	SchemeBuilder.Register(&Cloudtrail{}, &CloudtrailList{})
}
