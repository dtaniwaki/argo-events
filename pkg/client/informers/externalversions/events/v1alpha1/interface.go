/*
Copyright 2018 The Argo Project Authors.

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
	internalinterfaces "github.com/argoproj/argo-events/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Events returns a EventInformer.
	Events() EventInformer
	// Sensors returns a SensorInformer.
	Sensors() SensorInformer
	// Triggers returns a TriggerInformer.
	Triggers() TriggerInformer
	// Wires returns a WireInformer.
	Wires() WireInformer
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

// Events returns a EventInformer.
func (v *version) Events() EventInformer {
	return &eventInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Sensors returns a SensorInformer.
func (v *version) Sensors() SensorInformer {
	return &sensorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Triggers returns a TriggerInformer.
func (v *version) Triggers() TriggerInformer {
	return &triggerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Wires returns a WireInformer.
func (v *version) Wires() WireInformer {
	return &wireInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
