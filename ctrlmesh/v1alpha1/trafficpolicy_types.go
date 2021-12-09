/*
Copyright 2021 The Kruise Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TrafficPolicySpec defines the desired state of TrafficPolicy
type TrafficPolicySpec struct {
	TargetVirtualApps []TrafficTargetVirtualApp `json:"targetVirtualApps"`
	CircuitBreaking   *TrafficCircuitBreaking   `json:"circuitBreaking,omitempty"`
	RateLimiting      *TrafficRateLimiting      `json:"rateLimiting,omitempty"`
}

// TrafficTargetVirtualApp is the target VirtualApp and its optional specific subsets.
type TrafficTargetVirtualApp struct {
	Name            string   `json:"name"`
	SpecificSubsets []string `json:"specificSubsets,omitempty"`
}

// TrafficCircuitBreaking defines policies that ctrlmesh-proxy should intercept the requests.
type TrafficCircuitBreaking struct {
	APIServer *TrafficAPIServerRules `json:"apiServer,omitempty"`
	//Webhook   *TrafficWebhookRules   `json:"webhook,omitempty"`
}

type TrafficRateLimiting struct {
	RatePolicies []TrafficRateLimitingPolicy `json:"ratePolicies,omitempty"`
}

type TrafficRateLimitingPolicy struct {
	Rules TrafficAPIServerRules `json:"rules"`

	MaxInFlight        *int32                                 `json:"maxInFlight,omitempty"`
	Bucket             *TrafficRateLimitingBucket             `json:"bucket,omitempty"`
	ExponentialBackoff *TrafficRateLimitingExponentialBackoff `json:"exponentialBackoff,omitempty"`
}

type TrafficRateLimitingBucket struct {
	QPS   int32 `json:"qps"`
	Burst int32 `json:"burst"`
}

type TrafficRateLimitingExponentialBackoff struct {
	BaseDelayInMillisecond   int32 `json:"baseDelayInMillisecond"`
	MaxDelayInMillisecond    int32 `json:"maxDelayInMillisecond"`
	ContinuouslyFailureTimes int32 `json:"continuouslyFailureTimes,omitempty"`
}

// TrafficAPIServerRules contains rules for apiserver requests.
type TrafficAPIServerRules struct {
	// `resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the
	// target resource.
	// At least one of `resourceRules` and `nonResourceRules` has to be non-empty.
	ResourceRules []ResourcePolicyRule `json:"resourceRules,omitempty"`
	// `nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb
	// and the target non-resource URL.
	NonResourceRules []NonResourcePolicyRule `json:"nonResourceRules,omitempty"`
}

// ResourcePolicyRule is a predicate that matches some resource
// requests, testing the request's verb and the target resource. A
// ResourcePolicyRule matches a resource request if and only if: (a)
// at least one member of verbs matches the request, (b) at least one
// member of apiGroups matches the request, (c) at least one member of
// resources matches the request, and (d) least one member of
// namespaces matches the request.
type ResourcePolicyRule struct {
	// `verbs` is a list of matching verbs and may not be empty.
	// "*" matches all verbs and, if present, must be the only entry.
	// +listType=set
	// Required.
	Verbs []string `json:"verbs" protobuf:"bytes,1,rep,name=verbs"`

	// `apiGroups` is a list of matching API groups and may not be empty.
	// "*" matches all API groups and, if present, must be the only entry.
	// +listType=set
	// Required.
	APIGroups []string `json:"apiGroups" protobuf:"bytes,2,rep,name=apiGroups"`

	// `resources` is a list of matching resources (i.e., lowercase
	// and plural) with, if desired, subresource.  For example, [
	// "services", "nodes/status" ].  This list may not be empty.
	// "*" matches all resources and, if present, must be the only entry.
	// Required.
	// +listType=set
	Resources []string `json:"resources" protobuf:"bytes,3,rep,name=resources"`

	// `clusterScope` indicates whether to match requests that do not
	// specify a namespace (which happens either because the resource
	// is not namespaced or the request targets all namespaces).
	// If this field is omitted or false then the `namespaces` field
	// must contain a non-empty list.
	// +optional
	ClusterScope bool `json:"clusterScope,omitempty" protobuf:"varint,4,opt,name=clusterScope"`

	// `namespaces` is a list of target namespaces that restricts
	// matches.  A request that specifies a target namespace matches
	// only if either (a) this list contains that target namespace or
	// (b) this list contains "*".  Note that "*" matches any
	// specified namespace but does not match a request that _does
	// not specify_ a namespace (see the `clusterScope` field for
	// that).
	// This list may be empty, but only if `clusterScope` is true.
	// +optional
	// +listType=set
	Namespaces []string `json:"namespaces" protobuf:"bytes,5,rep,name=namespaces"`
}

// NonResourcePolicyRule is a predicate that matches non-resource requests according to their verb and the
// target non-resource URL. A NonResourcePolicyRule matches a request if and only if both (a) at least one member
// of verbs matches the request and (b) at least one member of nonResourceURLs matches the request.
type NonResourcePolicyRule struct {
	// `verbs` is a list of matching verbs and may not be empty.
	// "*" matches all verbs. If it is present, it must be the only entry.
	// +listType=set
	// Required.
	Verbs []string `json:"verbs" protobuf:"bytes,1,rep,name=verbs"`
	// `nonResourceURLs` is a set of url prefixes that a user should have access to and may not be empty.
	// For example:
	//   - "/healthz" is legal
	//   - "/hea*" is illegal
	//   - "/hea" is legal but matches nothing
	//   - "/hea/*" also matches nothing
	//   - "/healthz/*" matches all per-component health checks.
	// "*" matches all non-resource urls. if it is present, it must be the only entry.
	// +listType=set
	// Required.
	NonResourceURLs []string `json:"nonResourceURLs" protobuf:"bytes,6,rep,name=nonResourceURLs"`
}

// TrafficPolicyStatus defines the observed state of TrafficPolicy
type TrafficPolicyStatus struct {
}

// +genclient
// +k8s:openapi-gen=true
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TrafficPolicy is the Schema for the trafficpolicies API
type TrafficPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficPolicySpec   `json:"spec,omitempty"`
	Status TrafficPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TrafficPolicyList contains a list of TrafficPolicy
type TrafficPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficPolicy{}, &TrafficPolicyList{})
}
