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

package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type PackageSrcConfig struct {
	// +optional
	// tar.gz
	Type string `json:"type"`

	// +optional
	// untar path on dst node
	DstPath string `json:"dstpath"`

	// +optional
	ArmSrc string `json:"armsrc"`

	// +optional
	X86Src string `json:"x86src"`
}

type PackageConfig struct {
	Name string `json:"name"`
	// repo bin file dir image json shell
	Type     string `json:"type"`
	Dst      string `json:"dst,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	TimeOut  string `json:"timeout,omitempty"`
}

type AdditionConfig struct {
	// +optional
	Master []*PackageConfig `json:"master"`

	// +optional
	Worker []*PackageConfig `json:"worker"`

	// +optional
	ETCD []*PackageConfig `json:"etcd"`

	// +optional
	LoadBalance []*PackageConfig `json:"loadbalance"`
}

type InstallConfig struct {
	PackageSrc *PackageSrcConfig `json:"package-source"`

	// +optional
	KubernetesMaster []*PackageConfig `json:"kubernetes-master"`

	// +optional
	KubernetesWorker []*PackageConfig `json:"kubernetes-worker"`

	// +optional
	Network []*PackageConfig `json:"network"`

	// +optional
	ETCD []*PackageConfig `json:"etcd"`

	// +optional
	LoadBalance []*PackageConfig `json:"loadbalance"`

	// +optional
	Container []*PackageConfig `json:"container"`

	// +optional
	Image []*PackageConfig `json:"image"`

	// +optional
	Addition AdditionConfig `json:"addition"`
}

type OpenPorts struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	Port *int32 `json:"port"`

	// tcp/udp
	Protocol string `json:"protocol"`
}
type OpenPortsConfig struct {
	// +optional
	Master []*OpenPorts `json:"master"`

	// +optional
	Worker []*OpenPorts `json:"worker"`

	// +optional
	ETCD []*OpenPorts `json:"etcd"`

	// +optional
	LoadBalance []*OpenPorts `json:"loadbalance"`
}

type ClusterNetworkConfig struct {
	// config for cluster service network
	ServiceCidr    string `json:"service-cidr"`
	ServiceDnsIp   string `json:"service-dns-ip"`
	ServiceGateway string `json:"service-gateway"`

	// config for network of pod
	PodCidr   string `json:"pod-cidr"`
	PodPlugin string `json:"pod-plugin,omitempty"`
	// +optional
	PodPluginArgs map[string]string `json:"pod-plugin-args,omitempty"`
}

type APIEndpointConfig struct {
	Advertise string `json:"advertise,omitempty"`
	//+kubebuilder:validation:Minimum=0
	//+kubebuilder:validation:Maximum=65535
	BindPort *int32 `json:"bind-port,omitempty"`
}

type RuntimeConfig struct {
	Runtime         string `json:"runtime,omitempty"`
	RuntimeEndpoint string `json:"runtime-endpoint,omitempty"`
}

type RequireMachineConfig struct {
	Number int32 `json:"number"`

	// require machie need in which cidr
	// +optional
	Features map[string]string `json:"features,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// machines for master nodes
	//+kubebuilder:validation:Required
	MasterRequire RequireMachineConfig `json:"masterRequire"`

	// machines for worker nodes
	//+kubebuilder:validation:Required
	WorkerRequire RequireMachineConfig `json:"workerRequire"`

	// machines for etcd nodes
	EtcdRequire RequireMachineConfig `json:"etcdRequire,omitempty"`

	// machines for loadbalance
	LoadbalanceRequires RequireMachineConfig `json:"loadbalanceRequires,omitempty"`
	LoadbalanceBindPort int32                `json:"loadbalance-bindport,omitempty"`

	// MachineLoginSecret save user/password for ssh login
	//+kubebuilder:validation:Required
	MachineLoginSecret *v1.ObjectReference `json:"machineLoginSecret,omitempty"`

	// PackagePersistentVolumeClaim for pod
	//+kubebuilder:validation:Required
	PackagePersistentVolumeClaim *v1.ObjectReference `json:"packagePersistentVolumeClaim,omitempty"`

	InstallConfig InstallConfig `json:"install"`

	OpenPorts OpenPortsConfig `json:"open-ports"`

	ApiEndpoint APIEndpointConfig `json:"apiendpoint,omitempty"`

	Runtime RuntimeConfig `json:"runtime,omitempty"`

	// network config of cluster
	Network ClusterNetworkConfig `json:"network,omitempty"`

	Addons []string `json:"addons,omitempty"`
}

type JobHistory struct {
	Name       string       `json:"name"`
	StartTime  metav1.Time  `json:"start-time"`
	FinishTime *metav1.Time `json:"finish-time,omitempty"`
	Message    string       `json:"message,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	MachineLoginSecretRef *v1.ObjectReference `json:"machineLoginSecretRef,omitempty"`

	PackagePersistentVolumeClaimRef *v1.ObjectReference `json:"packagePersistentVolumeClaimRef,omitempty"`

	MachineBindingRef *v1.ObjectReference `json:"machineBindingRef,omitempty"`
	ConfigRef         *v1.ObjectReference `json:"configRef,omitempty"`
	JobRef            *v1.ObjectReference `json:"jobRef,omitempty"`
	JobHistorys       []*JobHistory       `json:"jobHistorys,omitempty"`

	HasCluster bool   `json:"hasCluster,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
	Message    string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

func (c *Cluster) IsCreated() bool {
	return c.Status.HasCluster
}

//+kubebuilder:object:root=true

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}