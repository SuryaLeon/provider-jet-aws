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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DelegationSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	NameServers []*string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`
}

type DelegationSetParameters struct {

	// +kubebuilder:validation:Optional
	ReferenceName *string `json:"referenceName,omitempty" tf:"reference_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// DelegationSetSpec defines the desired state of DelegationSet
type DelegationSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DelegationSetParameters `json:"forProvider"`
}

// DelegationSetStatus defines the observed state of DelegationSet.
type DelegationSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DelegationSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DelegationSet is the Schema for the DelegationSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DelegationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DelegationSetSpec   `json:"spec"`
	Status            DelegationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DelegationSetList contains a list of DelegationSets
type DelegationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DelegationSet `json:"items"`
}

// Repository type metadata.
var (
	DelegationSet_Kind             = "DelegationSet"
	DelegationSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DelegationSet_Kind}.String()
	DelegationSet_KindAPIVersion   = DelegationSet_Kind + "." + CRDGroupVersion.String()
	DelegationSet_GroupVersionKind = CRDGroupVersion.WithKind(DelegationSet_Kind)
)

func init() {
	SchemeBuilder.Register(&DelegationSet{}, &DelegationSetList{})
}
