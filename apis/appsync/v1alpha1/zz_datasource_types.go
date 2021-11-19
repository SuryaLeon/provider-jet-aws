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

type DatasourceObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type DatasourceParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DynamodbConfig []DynamodbConfigParameters `json:"dynamodbConfig,omitempty" tf:"dynamodb_config,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchConfig []ElasticsearchConfigParameters `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPConfig []HTTPConfigParameters `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DynamodbConfigObservation struct {
}

type DynamodbConfigParameters struct {

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`

	// +kubebuilder:validation:Optional
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials,omitempty"`
}

type ElasticsearchConfigObservation struct {
}

type ElasticsearchConfigParameters struct {

	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type HTTPConfigObservation struct {
}

type HTTPConfigParameters struct {

	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`
}

type LambdaConfigObservation struct {
}

type LambdaConfigParameters struct {

	// +kubebuilder:validation:Required
	FunctionArn *string `json:"functionArn" tf:"function_arn,omitempty"`
}

// DatasourceSpec defines the desired state of Datasource
type DatasourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasourceParameters `json:"forProvider"`
}

// DatasourceStatus defines the observed state of Datasource.
type DatasourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Datasource is the Schema for the Datasources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Datasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasourceSpec   `json:"spec"`
	Status            DatasourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasourceList contains a list of Datasources
type DatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datasource `json:"items"`
}

// Repository type metadata.
var (
	Datasource_Kind             = "Datasource"
	Datasource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datasource_Kind}.String()
	Datasource_KindAPIVersion   = Datasource_Kind + "." + CRDGroupVersion.String()
	Datasource_GroupVersionKind = CRDGroupVersion.WithKind(Datasource_Kind)
)

func init() {
	SchemeBuilder.Register(&Datasource{}, &DatasourceList{})
}
