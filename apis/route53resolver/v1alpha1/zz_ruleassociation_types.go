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

type RuleAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleAssociationParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResolverRuleID *string `json:"resolverRuleId" tf:"resolver_rule_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// RuleAssociationSpec defines the desired state of RuleAssociation
type RuleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleAssociationParameters `json:"forProvider"`
}

// RuleAssociationStatus defines the observed state of RuleAssociation.
type RuleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleAssociation is the Schema for the RuleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleAssociationSpec   `json:"spec"`
	Status            RuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleAssociationList contains a list of RuleAssociations
type RuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	RuleAssociation_Kind             = "RuleAssociation"
	RuleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleAssociation_Kind}.String()
	RuleAssociation_KindAPIVersion   = RuleAssociation_Kind + "." + CRDGroupVersion.String()
	RuleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(RuleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleAssociation{}, &RuleAssociationList{})
}
