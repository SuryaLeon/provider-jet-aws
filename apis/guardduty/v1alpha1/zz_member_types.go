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

type MemberObservation struct {
	RelationshipStatus *string `json:"relationshipStatus,omitempty" tf:"relationship_status,omitempty"`
}

type MemberParameters struct {

	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Required
	DetectorID *string `json:"detectorId" tf:"detector_id,omitempty"`

	// +kubebuilder:validation:Optional
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message,omitempty"`

	// +kubebuilder:validation:Optional
	Invite *bool `json:"invite,omitempty" tf:"invite,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberSpec   `json:"spec"`
	Status            MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
