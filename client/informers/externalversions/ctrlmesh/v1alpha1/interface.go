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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openkruise/controllermesh-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ManagerStates returns a ManagerStateInformer.
	ManagerStates() ManagerStateInformer
	// TrafficPolicies returns a TrafficPolicyInformer.
	TrafficPolicies() TrafficPolicyInformer
	// VirtualApps returns a VirtualAppInformer.
	VirtualApps() VirtualAppInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ManagerStates returns a ManagerStateInformer.
func (v *version) ManagerStates() ManagerStateInformer {
	return &managerStateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// TrafficPolicies returns a TrafficPolicyInformer.
func (v *version) TrafficPolicies() TrafficPolicyInformer {
	return &trafficPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualApps returns a VirtualAppInformer.
func (v *version) VirtualApps() VirtualAppInformer {
	return &virtualAppInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
