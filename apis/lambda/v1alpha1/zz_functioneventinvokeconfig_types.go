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

type DestinationConfigOnFailureObservation struct {
}

type DestinationConfigOnFailureParameters struct {

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`
}

type FunctionEventInvokeConfigDestinationConfigObservation struct {
}

type FunctionEventInvokeConfigDestinationConfigParameters struct {

	// +kubebuilder:validation:Optional
	OnFailure []DestinationConfigOnFailureParameters `json:"onFailure,omitempty" tf:"on_failure,omitempty"`

	// +kubebuilder:validation:Optional
	OnSuccess []OnSuccessParameters `json:"onSuccess,omitempty" tf:"on_success,omitempty"`
}

type FunctionEventInvokeConfigObservation struct {
}

type FunctionEventInvokeConfigParameters struct {

	// +kubebuilder:validation:Optional
	DestinationConfig []FunctionEventInvokeConfigDestinationConfigParameters `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumEventAgeInSeconds *int64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *int64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

type OnSuccessObservation struct {
}

type OnSuccessParameters struct {

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`
}

// FunctionEventInvokeConfigSpec defines the desired state of FunctionEventInvokeConfig
type FunctionEventInvokeConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionEventInvokeConfigParameters `json:"forProvider"`
}

// FunctionEventInvokeConfigStatus defines the observed state of FunctionEventInvokeConfig.
type FunctionEventInvokeConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionEventInvokeConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionEventInvokeConfig is the Schema for the FunctionEventInvokeConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FunctionEventInvokeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionEventInvokeConfigSpec   `json:"spec"`
	Status            FunctionEventInvokeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionEventInvokeConfigList contains a list of FunctionEventInvokeConfigs
type FunctionEventInvokeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionEventInvokeConfig `json:"items"`
}

// Repository type metadata.
var (
	FunctionEventInvokeConfig_Kind             = "FunctionEventInvokeConfig"
	FunctionEventInvokeConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionEventInvokeConfig_Kind}.String()
	FunctionEventInvokeConfig_KindAPIVersion   = FunctionEventInvokeConfig_Kind + "." + CRDGroupVersion.String()
	FunctionEventInvokeConfig_GroupVersionKind = CRDGroupVersion.WithKind(FunctionEventInvokeConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionEventInvokeConfig{}, &FunctionEventInvokeConfigList{})
}
