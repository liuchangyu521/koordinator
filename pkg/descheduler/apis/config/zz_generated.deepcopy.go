//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Koordinator Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultEvictorArgs) DeepCopyInto(out *DefaultEvictorArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.MaxNoOfPodsToEvictPerNode != nil {
		in, out := &in.MaxNoOfPodsToEvictPerNode, &out.MaxNoOfPodsToEvictPerNode
		*out = new(int)
		**out = **in
	}
	if in.MaxNoOfPodsToEvictPerNamespace != nil {
		in, out := &in.MaxNoOfPodsToEvictPerNamespace, &out.MaxNoOfPodsToEvictPerNamespace
		*out = new(int)
		**out = **in
	}
	if in.PriorityThreshold != nil {
		in, out := &in.PriorityThreshold, &out.PriorityThreshold
		*out = new(PriorityThreshold)
		(*in).DeepCopyInto(*out)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultEvictorArgs.
func (in *DefaultEvictorArgs) DeepCopy() *DefaultEvictorArgs {
	if in == nil {
		return nil
	}
	out := new(DefaultEvictorArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefaultEvictorArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulerConfiguration) DeepCopyInto(out *DeschedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.LeaderElection = in.LeaderElection
	out.ClientConnection = in.ClientConnection
	out.DebuggingConfiguration = in.DebuggingConfiguration
	out.DeschedulingInterval = in.DeschedulingInterval
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]DeschedulerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulerConfiguration.
func (in *DeschedulerConfiguration) DeepCopy() *DeschedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeschedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeschedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulerProfile) DeepCopyInto(out *DeschedulerProfile) {
	*out = *in
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(Plugins)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulerProfile.
func (in *DeschedulerProfile) DeepCopy() *DeschedulerProfile {
	if in == nil {
		return nil
	}
	out := new(DeschedulerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LowNodeLoadArgs) DeepCopyInto(out *LowNodeLoadArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.EvictableNamespaces != nil {
		in, out := &in.EvictableNamespaces, &out.EvictableNamespaces
		*out = new(Namespaces)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSelectors != nil {
		in, out := &in.PodSelectors, &out.PodSelectors
		*out = make([]LowNodeLoadPodSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HighThresholds != nil {
		in, out := &in.HighThresholds, &out.HighThresholds
		*out = make(ResourceThresholds, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LowThresholds != nil {
		in, out := &in.LowThresholds, &out.LowThresholds
		*out = make(ResourceThresholds, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LowNodeLoadArgs.
func (in *LowNodeLoadArgs) DeepCopy() *LowNodeLoadArgs {
	if in == nil {
		return nil
	}
	out := new(LowNodeLoadArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LowNodeLoadArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LowNodeLoadPodSelector) DeepCopyInto(out *LowNodeLoadPodSelector) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LowNodeLoadPodSelector.
func (in *LowNodeLoadPodSelector) DeepCopy() *LowNodeLoadPodSelector {
	if in == nil {
		return nil
	}
	out := new(LowNodeLoadPodSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationControllerArgs) DeepCopyInto(out *MigrationControllerArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = new(Namespaces)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxMigratingPerNode != nil {
		in, out := &in.MaxMigratingPerNode, &out.MaxMigratingPerNode
		*out = new(int32)
		**out = **in
	}
	if in.MaxMigratingPerNamespace != nil {
		in, out := &in.MaxMigratingPerNamespace, &out.MaxMigratingPerNamespace
		*out = new(int32)
		**out = **in
	}
	if in.MaxMigratingPerWorkload != nil {
		in, out := &in.MaxMigratingPerWorkload, &out.MaxMigratingPerWorkload
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailablePerWorkload != nil {
		in, out := &in.MaxUnavailablePerWorkload, &out.MaxUnavailablePerWorkload
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.ObjectLimiters != nil {
		in, out := &in.ObjectLimiters, &out.ObjectLimiters
		*out = make(ObjectLimiterMap, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.DefaultJobTTL = in.DefaultJobTTL
	if in.DefaultDeleteOptions != nil {
		in, out := &in.DefaultDeleteOptions, &out.DefaultDeleteOptions
		*out = new(v1.DeleteOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationControllerArgs.
func (in *MigrationControllerArgs) DeepCopy() *MigrationControllerArgs {
	if in == nil {
		return nil
	}
	out := new(MigrationControllerArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigrationControllerArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationObjectLimiter) DeepCopyInto(out *MigrationObjectLimiter) {
	*out = *in
	out.Duration = in.Duration
	if in.MaxMigrating != nil {
		in, out := &in.MaxMigrating, &out.MaxMigrating
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationObjectLimiter.
func (in *MigrationObjectLimiter) DeepCopy() *MigrationObjectLimiter {
	if in == nil {
		return nil
	}
	out := new(MigrationObjectLimiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespaces) DeepCopyInto(out *Namespaces) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespaces.
func (in *Namespaces) DeepCopy() *Namespaces {
	if in == nil {
		return nil
	}
	out := new(Namespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ObjectLimiterMap) DeepCopyInto(out *ObjectLimiterMap) {
	{
		in := &in
		*out = make(ObjectLimiterMap, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectLimiterMap.
func (in ObjectLimiterMap) DeepCopy() ObjectLimiterMap {
	if in == nil {
		return nil
	}
	out := new(ObjectLimiterMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	if in.Args != nil {
		out.Args = in.Args.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSet) DeepCopyInto(out *PluginSet) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSet.
func (in *PluginSet) DeepCopy() *PluginSet {
	if in == nil {
		return nil
	}
	out := new(PluginSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugins) DeepCopyInto(out *Plugins) {
	*out = *in
	in.Deschedule.DeepCopyInto(&out.Deschedule)
	in.Balance.DeepCopyInto(&out.Balance)
	in.Evictor.DeepCopyInto(&out.Evictor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugins.
func (in *Plugins) DeepCopy() *Plugins {
	if in == nil {
		return nil
	}
	out := new(Plugins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityThreshold) DeepCopyInto(out *PriorityThreshold) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityThreshold.
func (in *PriorityThreshold) DeepCopy() *PriorityThreshold {
	if in == nil {
		return nil
	}
	out := new(PriorityThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemovePodsViolatingNodeAffinityArgs) DeepCopyInto(out *RemovePodsViolatingNodeAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = new(Namespaces)
		(*in).DeepCopyInto(*out)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeAffinityType != nil {
		in, out := &in.NodeAffinityType, &out.NodeAffinityType
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemovePodsViolatingNodeAffinityArgs.
func (in *RemovePodsViolatingNodeAffinityArgs) DeepCopy() *RemovePodsViolatingNodeAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(RemovePodsViolatingNodeAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemovePodsViolatingNodeAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ResourceThresholds) DeepCopyInto(out *ResourceThresholds) {
	{
		in := &in
		*out = make(ResourceThresholds, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceThresholds.
func (in ResourceThresholds) DeepCopy() ResourceThresholds {
	if in == nil {
		return nil
	}
	out := new(ResourceThresholds)
	in.DeepCopyInto(out)
	return *out
}
