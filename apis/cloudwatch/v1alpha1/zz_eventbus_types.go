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

type EventBusObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type EventBusParameters struct {

	// +kubebuilder:validation:Optional
	EventSourceName *string `json:"eventSourceName,omitempty" tf:"event_source_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EventBusSpec defines the desired state of EventBus
type EventBusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventBusParameters `json:"forProvider"`
}

// EventBusStatus defines the observed state of EventBus.
type EventBusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventBusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventBus is the Schema for the EventBuss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type EventBus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventBusSpec   `json:"spec"`
	Status            EventBusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventBusList contains a list of EventBuss
type EventBusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventBus `json:"items"`
}

// Repository type metadata.
var (
	EventBus_Kind             = "EventBus"
	EventBus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventBus_Kind}.String()
	EventBus_KindAPIVersion   = EventBus_Kind + "." + CRDGroupVersion.String()
	EventBus_GroupVersionKind = CRDGroupVersion.WithKind(EventBus_Kind)
)

func init() {
	SchemeBuilder.Register(&EventBus{}, &EventBusList{})
}
