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

type DestinationConfigObservation struct {
}

type DestinationConfigParameters struct {

	// +kubebuilder:validation:Optional
	OnFailure []OnFailureParameters `json:"onFailure,omitempty" tf:"on_failure,omitempty"`
}

type EventSourceMappingObservation struct {
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	LastProcessingResult *string `json:"lastProcessingResult,omitempty" tf:"last_processing_result,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	StateTransitionReason *string `json:"stateTransitionReason,omitempty" tf:"state_transition_reason,omitempty"`

	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type EventSourceMappingParameters struct {

	// +kubebuilder:validation:Optional
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// +kubebuilder:validation:Optional
	BisectBatchOnFunctionError *bool `json:"bisectBatchOnFunctionError,omitempty" tf:"bisect_batch_on_function_error,omitempty"`

	// +kubebuilder:validation:Optional
	DestinationConfig []DestinationConfigParameters `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn,omitempty"`

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Optional
	FunctionResponseTypes []*string `json:"functionResponseTypes,omitempty" tf:"function_response_types,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumBatchingWindowInSeconds *int64 `json:"maximumBatchingWindowInSeconds,omitempty" tf:"maximum_batching_window_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumRecordAgeInSeconds *int64 `json:"maximumRecordAgeInSeconds,omitempty" tf:"maximum_record_age_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *int64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// +kubebuilder:validation:Optional
	ParallelizationFactor *int64 `json:"parallelizationFactor,omitempty" tf:"parallelization_factor,omitempty"`

	// +kubebuilder:validation:Optional
	Queues []*string `json:"queues,omitempty" tf:"queues,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SelfManagedEventSource []SelfManagedEventSourceParameters `json:"selfManagedEventSource,omitempty" tf:"self_managed_event_source,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAccessConfiguration []SourceAccessConfigurationParameters `json:"sourceAccessConfiguration,omitempty" tf:"source_access_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`

	// +kubebuilder:validation:Optional
	StartingPositionTimestamp *string `json:"startingPositionTimestamp,omitempty" tf:"starting_position_timestamp,omitempty"`

	// +kubebuilder:validation:Optional
	Topics []*string `json:"topics,omitempty" tf:"topics,omitempty"`

	// +kubebuilder:validation:Optional
	TumblingWindowInSeconds *int64 `json:"tumblingWindowInSeconds,omitempty" tf:"tumbling_window_in_seconds,omitempty"`
}

type OnFailureObservation struct {
}

type OnFailureParameters struct {

	// +kubebuilder:validation:Required
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`
}

type SelfManagedEventSourceObservation struct {
}

type SelfManagedEventSourceParameters struct {

	// +kubebuilder:validation:Required
	Endpoints map[string]*string `json:"endpoints" tf:"endpoints,omitempty"`
}

type SourceAccessConfigurationObservation struct {
}

type SourceAccessConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

// EventSourceMappingSpec defines the desired state of EventSourceMapping
type EventSourceMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSourceMappingParameters `json:"forProvider"`
}

// EventSourceMappingStatus defines the observed state of EventSourceMapping.
type EventSourceMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSourceMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceMapping is the Schema for the EventSourceMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type EventSourceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSourceMappingSpec   `json:"spec"`
	Status            EventSourceMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceMappingList contains a list of EventSourceMappings
type EventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSourceMapping `json:"items"`
}

// Repository type metadata.
var (
	EventSourceMapping_Kind             = "EventSourceMapping"
	EventSourceMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSourceMapping_Kind}.String()
	EventSourceMapping_KindAPIVersion   = EventSourceMapping_Kind + "." + CRDGroupVersion.String()
	EventSourceMapping_GroupVersionKind = CRDGroupVersion.WithKind(EventSourceMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSourceMapping{}, &EventSourceMappingList{})
}
