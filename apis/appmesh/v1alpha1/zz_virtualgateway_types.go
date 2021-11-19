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

type AccessLogFileObservation struct {
}

type AccessLogFileParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type AccessLogObservation struct {
}

type AccessLogParameters struct {

	// +kubebuilder:validation:Optional
	File []AccessLogFileParameters `json:"file,omitempty" tf:"file,omitempty"`
}

type AcmObservation struct {
}

type AcmParameters struct {

	// +kubebuilder:validation:Required
	CertificateAuthorityArns []*string `json:"certificateAuthorityArns" tf:"certificate_authority_arns,omitempty"`
}

type BackendDefaultsObservation struct {
}

type BackendDefaultsParameters struct {

	// +kubebuilder:validation:Optional
	ClientPolicy []ClientPolicyParameters `json:"clientPolicy,omitempty" tf:"client_policy,omitempty"`
}

type CertificateAcmObservation struct {
}

type CertificateAcmParameters struct {

	// +kubebuilder:validation:Required
	CertificateArn *string `json:"certificateArn" tf:"certificate_arn,omitempty"`
}

type CertificateFileObservation struct {
}

type CertificateFileParameters struct {

	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`

	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type CertificateObservation struct {
}

type CertificateParameters struct {

	// +kubebuilder:validation:Optional
	File []FileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// +kubebuilder:validation:Optional
	Sds []SdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type CertificateSdsObservation struct {
}

type CertificateSdsParameters struct {

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type ClientPolicyObservation struct {
}

type ClientPolicyParameters struct {

	// +kubebuilder:validation:Optional
	TLS []TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ConnectionPoolObservation struct {
}

type ConnectionPoolParameters struct {

	// +kubebuilder:validation:Optional
	Grpc []GrpcParameters `json:"grpc,omitempty" tf:"grpc,omitempty"`

	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// +kubebuilder:validation:Optional
	Http2 []Http2Parameters `json:"http2,omitempty" tf:"http2,omitempty"`
}

type FileObservation struct {
}

type FileParameters struct {

	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`

	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type GrpcObservation struct {
}

type GrpcParameters struct {

	// +kubebuilder:validation:Required
	MaxRequests *int64 `json:"maxRequests" tf:"max_requests,omitempty"`
}

type HTTPObservation struct {
}

type HTTPParameters struct {

	// +kubebuilder:validation:Required
	MaxConnections *int64 `json:"maxConnections" tf:"max_connections,omitempty"`

	// +kubebuilder:validation:Optional
	MaxPendingRequests *int64 `json:"maxPendingRequests,omitempty" tf:"max_pending_requests,omitempty"`
}

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// +kubebuilder:validation:Required
	HealthyThreshold *int64 `json:"healthyThreshold" tf:"healthy_threshold,omitempty"`

	// +kubebuilder:validation:Required
	IntervalMillis *int64 `json:"intervalMillis" tf:"interval_millis,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	TimeoutMillis *int64 `json:"timeoutMillis" tf:"timeout_millis,omitempty"`

	// +kubebuilder:validation:Required
	UnhealthyThreshold *int64 `json:"unhealthyThreshold" tf:"unhealthy_threshold,omitempty"`
}

type Http2Observation struct {
}

type Http2Parameters struct {

	// +kubebuilder:validation:Required
	MaxRequests *int64 `json:"maxRequests" tf:"max_requests,omitempty"`
}

type ListenerObservation struct {
}

type ListenerParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionPool []ConnectionPoolParameters `json:"connectionPool,omitempty" tf:"connection_pool,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// +kubebuilder:validation:Required
	PortMapping []PortMappingParameters `json:"portMapping" tf:"port_mapping,omitempty"`

	// +kubebuilder:validation:Optional
	TLS []ListenerTLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ListenerTLSObservation struct {
}

type ListenerTLSParameters struct {

	// +kubebuilder:validation:Required
	Certificate []TLSCertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Validation []TLSValidationParameters `json:"validation,omitempty" tf:"validation,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// +kubebuilder:validation:Optional
	AccessLog []AccessLogParameters `json:"accessLog,omitempty" tf:"access_log,omitempty"`
}

type PortMappingObservation struct {
}

type PortMappingParameters struct {

	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type SdsObservation struct {
}

type SdsParameters struct {

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type SubjectAlternativeNamesMatchObservation struct {
}

type SubjectAlternativeNamesMatchParameters struct {

	// +kubebuilder:validation:Required
	Exact []*string `json:"exact" tf:"exact,omitempty"`
}

type SubjectAlternativeNamesObservation struct {
}

type SubjectAlternativeNamesParameters struct {

	// +kubebuilder:validation:Required
	Match []SubjectAlternativeNamesMatchParameters `json:"match" tf:"match,omitempty"`
}

type TLSCertificateObservation struct {
}

type TLSCertificateParameters struct {

	// +kubebuilder:validation:Optional
	Acm []CertificateAcmParameters `json:"acm,omitempty" tf:"acm,omitempty"`

	// +kubebuilder:validation:Optional
	File []CertificateFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// +kubebuilder:validation:Optional
	Sds []CertificateSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type TLSObservation struct {
}

type TLSParameters struct {

	// +kubebuilder:validation:Optional
	Certificate []CertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	Enforce *bool `json:"enforce,omitempty" tf:"enforce,omitempty"`

	// +kubebuilder:validation:Optional
	Ports []*int64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// +kubebuilder:validation:Required
	Validation []ValidationParameters `json:"validation" tf:"validation,omitempty"`
}

type TLSValidationObservation struct {
}

type TLSValidationParameters struct {

	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []ValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// +kubebuilder:validation:Required
	Trust []ValidationTrustParameters `json:"trust" tf:"trust,omitempty"`
}

type TrustFileObservation struct {
}

type TrustFileParameters struct {

	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`
}

type TrustObservation struct {
}

type TrustParameters struct {

	// +kubebuilder:validation:Optional
	Acm []AcmParameters `json:"acm,omitempty" tf:"acm,omitempty"`

	// +kubebuilder:validation:Optional
	File []TrustFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// +kubebuilder:validation:Optional
	Sds []TrustSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type TrustSdsObservation struct {
}

type TrustSdsParameters struct {

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type ValidationObservation struct {
}

type ValidationParameters struct {

	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []SubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// +kubebuilder:validation:Required
	Trust []TrustParameters `json:"trust" tf:"trust,omitempty"`
}

type ValidationSubjectAlternativeNamesMatchObservation struct {
}

type ValidationSubjectAlternativeNamesMatchParameters struct {

	// +kubebuilder:validation:Required
	Exact []*string `json:"exact" tf:"exact,omitempty"`
}

type ValidationSubjectAlternativeNamesObservation struct {
}

type ValidationSubjectAlternativeNamesParameters struct {

	// +kubebuilder:validation:Required
	Match []ValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match,omitempty"`
}

type ValidationTrustFileObservation struct {
}

type ValidationTrustFileParameters struct {

	// +kubebuilder:validation:Required
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`
}

type ValidationTrustObservation struct {
}

type ValidationTrustParameters struct {

	// +kubebuilder:validation:Optional
	File []ValidationTrustFileParameters `json:"file,omitempty" tf:"file,omitempty"`

	// +kubebuilder:validation:Optional
	Sds []ValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds,omitempty"`
}

type ValidationTrustSdsObservation struct {
}

type ValidationTrustSdsParameters struct {

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type VirtualGatewayObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VirtualGatewayParameters struct {

	// +kubebuilder:validation:Required
	MeshName *string `json:"meshName" tf:"mesh_name,omitempty"`

	// +kubebuilder:validation:Optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Spec []VirtualGatewaySpecParameters `json:"spec" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualGatewaySpecObservation struct {
}

type VirtualGatewaySpecParameters struct {

	// +kubebuilder:validation:Optional
	BackendDefaults []BackendDefaultsParameters `json:"backendDefaults,omitempty" tf:"backend_defaults,omitempty"`

	// +kubebuilder:validation:Required
	Listener []ListenerParameters `json:"listener" tf:"listener,omitempty"`

	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`
}

// VirtualGatewaySpec defines the desired state of VirtualGateway
type VirtualGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualGatewayParameters `json:"forProvider"`
}

// VirtualGatewayStatus defines the observed state of VirtualGateway.
type VirtualGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualGateway is the Schema for the VirtualGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type VirtualGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualGatewaySpec   `json:"spec"`
	Status            VirtualGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualGatewayList contains a list of VirtualGateways
type VirtualGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualGateway `json:"items"`
}

// Repository type metadata.
var (
	VirtualGateway_Kind             = "VirtualGateway"
	VirtualGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualGateway_Kind}.String()
	VirtualGateway_KindAPIVersion   = VirtualGateway_Kind + "." + CRDGroupVersion.String()
	VirtualGateway_GroupVersionKind = CRDGroupVersion.WithKind(VirtualGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualGateway{}, &VirtualGatewayList{})
}
