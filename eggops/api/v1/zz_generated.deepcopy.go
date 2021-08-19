// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpointConfig) DeepCopyInto(out *APIEndpointConfig) {
	*out = *in
	if in.BindPort != nil {
		in, out := &in.BindPort, &out.BindPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpointConfig.
func (in *APIEndpointConfig) DeepCopy() *APIEndpointConfig {
	if in == nil {
		return nil
	}
	out := new(APIEndpointConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionConfig) DeepCopyInto(out *AdditionConfig) {
	*out = *in
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.ETCD != nil {
		in, out := &in.ETCD, &out.ETCD
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.LoadBalance != nil {
		in, out := &in.LoadBalance, &out.LoadBalance
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionConfig.
func (in *AdditionConfig) DeepCopy() *AdditionConfig {
	if in == nil {
		return nil
	}
	out := new(AdditionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterNetworkConfig) DeepCopyInto(out *ClusterNetworkConfig) {
	*out = *in
	if in.PodPluginArgs != nil {
		in, out := &in.PodPluginArgs, &out.PodPluginArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterNetworkConfig.
func (in *ClusterNetworkConfig) DeepCopy() *ClusterNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.MasterRequire.DeepCopyInto(&out.MasterRequire)
	in.WorkerRequire.DeepCopyInto(&out.WorkerRequire)
	in.LoadbalanceRequires.DeepCopyInto(&out.LoadbalanceRequires)
	if in.MachineLoginSecret != nil {
		in, out := &in.MachineLoginSecret, &out.MachineLoginSecret
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Infrastructure != nil {
		in, out := &in.Infrastructure, &out.Infrastructure
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	in.ApiEndpoint.DeepCopyInto(&out.ApiEndpoint)
	out.Runtime = in.Runtime
	in.Network.DeepCopyInto(&out.Network)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.MachineLoginSecretRef != nil {
		in, out := &in.MachineLoginSecretRef, &out.MachineLoginSecretRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.PackagePersistentVolumeClaimRef != nil {
		in, out := &in.PackagePersistentVolumeClaimRef, &out.PackagePersistentVolumeClaimRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.MachineBindingRef != nil {
		in, out := &in.MachineBindingRef, &out.MachineBindingRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.ConfigRef != nil {
		in, out := &in.ConfigRef, &out.ConfigRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.JobRef != nil {
		in, out := &in.JobRef, &out.JobRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.JobHistorys != nil {
		in, out := &in.JobHistorys, &out.JobHistorys
		*out = make([]*JobHistory, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(JobHistory)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infrastructure) DeepCopyInto(out *Infrastructure) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infrastructure.
func (in *Infrastructure) DeepCopy() *Infrastructure {
	if in == nil {
		return nil
	}
	out := new(Infrastructure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Infrastructure) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureList) DeepCopyInto(out *InfrastructureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Infrastructure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureList.
func (in *InfrastructureList) DeepCopy() *InfrastructureList {
	if in == nil {
		return nil
	}
	out := new(InfrastructureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureSpec) DeepCopyInto(out *InfrastructureSpec) {
	*out = *in
	if in.PackagePersistentVolumeClaim != nil {
		in, out := &in.PackagePersistentVolumeClaim, &out.PackagePersistentVolumeClaim
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	in.InstallConfig.DeepCopyInto(&out.InstallConfig)
	in.OpenPorts.DeepCopyInto(&out.OpenPorts)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureSpec.
func (in *InfrastructureSpec) DeepCopy() *InfrastructureSpec {
	if in == nil {
		return nil
	}
	out := new(InfrastructureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureStatus) DeepCopyInto(out *InfrastructureStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureStatus.
func (in *InfrastructureStatus) DeepCopy() *InfrastructureStatus {
	if in == nil {
		return nil
	}
	out := new(InfrastructureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallConfig) DeepCopyInto(out *InstallConfig) {
	*out = *in
	if in.PackageSrc != nil {
		in, out := &in.PackageSrc, &out.PackageSrc
		*out = new(PackageSrcConfig)
		**out = **in
	}
	if in.KubernetesMaster != nil {
		in, out := &in.KubernetesMaster, &out.KubernetesMaster
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.KubernetesWorker != nil {
		in, out := &in.KubernetesWorker, &out.KubernetesWorker
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.ETCD != nil {
		in, out := &in.ETCD, &out.ETCD
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.LoadBalance != nil {
		in, out := &in.LoadBalance, &out.LoadBalance
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = make([]*PackageConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageConfig)
				**out = **in
			}
		}
	}
	in.Addition.DeepCopyInto(&out.Addition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallConfig.
func (in *InstallConfig) DeepCopy() *InstallConfig {
	if in == nil {
		return nil
	}
	out := new(InstallConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobHistory) DeepCopyInto(out *JobHistory) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobHistory.
func (in *JobHistory) DeepCopy() *JobHistory {
	if in == nil {
		return nil
	}
	out := new(JobHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Machine) DeepCopyInto(out *Machine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Machine.
func (in *Machine) DeepCopy() *Machine {
	if in == nil {
		return nil
	}
	out := new(Machine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Machine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineBinding) DeepCopyInto(out *MachineBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineBinding.
func (in *MachineBinding) DeepCopy() *MachineBinding {
	if in == nil {
		return nil
	}
	out := new(MachineBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineBindingList) DeepCopyInto(out *MachineBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineBindingList.
func (in *MachineBindingList) DeepCopy() *MachineBindingList {
	if in == nil {
		return nil
	}
	out := new(MachineBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineBindingSpec) DeepCopyInto(out *MachineBindingSpec) {
	*out = *in
	if in.MachineSets != nil {
		in, out := &in.MachineSets, &out.MachineSets
		*out = make([]MachineSetOfUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineBindingSpec.
func (in *MachineBindingSpec) DeepCopy() *MachineBindingSpec {
	if in == nil {
		return nil
	}
	out := new(MachineBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineBindingStatus) DeepCopyInto(out *MachineBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]MachineCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineBindingStatus.
func (in *MachineBindingStatus) DeepCopy() *MachineBindingStatus {
	if in == nil {
		return nil
	}
	out := new(MachineBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineCondition) DeepCopyInto(out *MachineCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineCondition.
func (in *MachineCondition) DeepCopy() *MachineCondition {
	if in == nil {
		return nil
	}
	out := new(MachineCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineList) DeepCopyInto(out *MachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Machine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineList.
func (in *MachineList) DeepCopy() *MachineList {
	if in == nil {
		return nil
	}
	out := new(MachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetOfUsage) DeepCopyInto(out *MachineSetOfUsage) {
	*out = *in
	if in.Machines != nil {
		in, out := &in.Machines, &out.Machines
		*out = make([]*Machine, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Machine)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetOfUsage.
func (in *MachineSetOfUsage) DeepCopy() *MachineSetOfUsage {
	if in == nil {
		return nil
	}
	out := new(MachineSetOfUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPorts) DeepCopyInto(out *OpenPorts) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPorts.
func (in *OpenPorts) DeepCopy() *OpenPorts {
	if in == nil {
		return nil
	}
	out := new(OpenPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenPortsConfig) DeepCopyInto(out *OpenPortsConfig) {
	*out = *in
	if in.Master != nil {
		in, out := &in.Master, &out.Master
		*out = make([]*OpenPorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OpenPorts)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = make([]*OpenPorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OpenPorts)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ETCD != nil {
		in, out := &in.ETCD, &out.ETCD
		*out = make([]*OpenPorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OpenPorts)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.LoadBalance != nil {
		in, out := &in.LoadBalance, &out.LoadBalance
		*out = make([]*OpenPorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OpenPorts)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenPortsConfig.
func (in *OpenPortsConfig) DeepCopy() *OpenPortsConfig {
	if in == nil {
		return nil
	}
	out := new(OpenPortsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageConfig) DeepCopyInto(out *PackageConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageConfig.
func (in *PackageConfig) DeepCopy() *PackageConfig {
	if in == nil {
		return nil
	}
	out := new(PackageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageSrcConfig) DeepCopyInto(out *PackageSrcConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageSrcConfig.
func (in *PackageSrcConfig) DeepCopy() *PackageSrcConfig {
	if in == nil {
		return nil
	}
	out := new(PackageSrcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequireMachineConfig) DeepCopyInto(out *RequireMachineConfig) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequireMachineConfig.
func (in *RequireMachineConfig) DeepCopy() *RequireMachineConfig {
	if in == nil {
		return nil
	}
	out := new(RequireMachineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeConfig) DeepCopyInto(out *RuntimeConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeConfig.
func (in *RuntimeConfig) DeepCopy() *RuntimeConfig {
	if in == nil {
		return nil
	}
	out := new(RuntimeConfig)
	in.DeepCopyInto(out)
	return out
}
