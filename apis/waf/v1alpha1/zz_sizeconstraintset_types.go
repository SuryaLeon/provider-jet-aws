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

type SizeConstraintSetObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type SizeConstraintSetParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SizeConstraints []SizeConstraintsParameters `json:"sizeConstraints,omitempty" tf:"size_constraints,omitempty"`
}

type SizeConstraintsFieldToMatchObservation struct {
}

type SizeConstraintsFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SizeConstraintsObservation struct {
}

type SizeConstraintsParameters struct {

	// +kubebuilder:validation:Required
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// +kubebuilder:validation:Required
	FieldToMatch []SizeConstraintsFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match,omitempty"`

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Required
	TextTransformation *string `json:"textTransformation" tf:"text_transformation,omitempty"`
}

// SizeConstraintSetSpec defines the desired state of SizeConstraintSet
type SizeConstraintSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SizeConstraintSetParameters `json:"forProvider"`
}

// SizeConstraintSetStatus defines the observed state of SizeConstraintSet.
type SizeConstraintSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SizeConstraintSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SizeConstraintSet is the Schema for the SizeConstraintSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SizeConstraintSetSpec   `json:"spec"`
	Status            SizeConstraintSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SizeConstraintSetList contains a list of SizeConstraintSets
type SizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SizeConstraintSet `json:"items"`
}

// Repository type metadata.
var (
	SizeConstraintSet_Kind             = "SizeConstraintSet"
	SizeConstraintSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SizeConstraintSet_Kind}.String()
	SizeConstraintSet_KindAPIVersion   = SizeConstraintSet_Kind + "." + CRDGroupVersion.String()
	SizeConstraintSet_GroupVersionKind = CRDGroupVersion.WithKind(SizeConstraintSet_Kind)
)

func init() {
	SchemeBuilder.Register(&SizeConstraintSet{}, &SizeConstraintSetList{})
}
