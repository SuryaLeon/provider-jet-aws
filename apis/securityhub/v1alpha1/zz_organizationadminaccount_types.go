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

type OrganizationAdminAccountObservation struct {
}

type OrganizationAdminAccountParameters struct {

	// +kubebuilder:validation:Required
	AdminAccountID *string `json:"adminAccountId" tf:"admin_account_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// OrganizationAdminAccountSpec defines the desired state of OrganizationAdminAccount
type OrganizationAdminAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationAdminAccountParameters `json:"forProvider"`
}

// OrganizationAdminAccountStatus defines the observed state of OrganizationAdminAccount.
type OrganizationAdminAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationAdminAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationAdminAccount is the Schema for the OrganizationAdminAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type OrganizationAdminAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationAdminAccountSpec   `json:"spec"`
	Status            OrganizationAdminAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationAdminAccountList contains a list of OrganizationAdminAccounts
type OrganizationAdminAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationAdminAccount `json:"items"`
}

// Repository type metadata.
var (
	OrganizationAdminAccount_Kind             = "OrganizationAdminAccount"
	OrganizationAdminAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationAdminAccount_Kind}.String()
	OrganizationAdminAccount_KindAPIVersion   = OrganizationAdminAccount_Kind + "." + CRDGroupVersion.String()
	OrganizationAdminAccount_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationAdminAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationAdminAccount{}, &OrganizationAdminAccountList{})
}
