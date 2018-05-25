package models

import (
	"k8s.io/api/core/v1"
)

type Pod struct {
	Name       string      `json:"name"`
	Namespace  string      `json:"namespace"`
	Desc       string      `json:"desc"`
	Containers []Container `json:"containers"`
}
type Container struct {
	Name          string      `json:"name"`
	Image         string      `json:"image"`
	ContainerPort int32       `json:"containerPort" protobuf:"varint,3,opt,name=containerPort"`
	Protocol      string      `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol"`
	Env           []v1.EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
}
type Deployment struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	Desc       string            `json:"desc"`
	Replicas   int32             `json:"replicas"`
	Containers []Container       `json:"containers"`
	// AutoScale AutoScale `json:"autoScale"
}
type ServiceType string
type Service struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Ports     []Port            `json:"ports"`
	Selector  map[string]string `json:"selector,omitempty" protobuf:"bytes,2,rep,name=selector"`
	Type      v1.ServiceType    `json:"type,omitempty" protobuf:"bytes,4,opt,name=type,casttype=ServiceType"`
}
type AutoScale struct {
	MinReplicas            int32 `json:"minReplicas"`
	MaxReplicas            int32 `json:"maxReplicas"`
	TargetCPUPercentage    int32 `json:"targetCPUPercentage"`
	TargetMemoryPercentage int32 `json:"targetMemoryPercentage"`
}
type Port struct {
	ContainerPort int    `json:"containerPort"`
	ServicePort   int32  `json:"servicePort"`
	Protocol      string `json:"protocol"`
	NodePort      int32  `json:"nodePort,omitempty" protobuf:"varint,5,opt,name=nodePort"`
}
type PodList struct {
	TypeMeta TypeMeta `json:"typeMeta,inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	ListMeta ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of pods.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md
	Items []Pod `json:"items" protobuf:"bytes,2,rep,name=items"`
}
type StatefulSetList struct {
	TypeMeta TypeMeta `json:",inline"`
	// +optional
	ListMeta ListMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items    []Statefulset `json:"items" protobuf:"bytes,2,rep,name=items"`
}
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	// In CamelCase.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}
type ListMeta struct {
	// selfLink is a URL representing this object.
	// Populated by the system.
	// Read-only.
	// +optional
	SelfLink string `json:"selfLink,omitempty" protobuf:"bytes,1,opt,name=selfLink"`

	// String that identifies the server's internal version of this object that
	// can be used by clients to determine when objects have changed.
	// Value must be treated as opaque by clients and passed unmodified back to the server.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,2,opt,name=resourceVersion"`

	// continue may be set if the user set a limit on the number of items returned, and indicates that
	// the server has more data available. The value is opaque and may be used to issue another request
	// to the endpoint that served this list to retrieve the next set of available objects. Continuing a
	// list may not be possible if the server configuration has changed or more than a few minutes have
	// passed. The resourceVersion field returned when using this continue value will be identical to
	// the value in the first response.
	Continue string `json:"continue,omitempty" protobuf:"bytes,3,opt,name=continue"`
}
type DeploymentList struct {
	TypeMeta TypeMeta `json:"typeMeta,inline"`
	// Standard list metadata.
	// +optional
	ListMeta ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of Deployments.
	Items []Deployment `json:"items" protobuf:"bytes,2,rep,name=items"`
}
type ServiceList struct {
	TypeMeta TypeMeta `json:"typeMeta,inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	ListMeta ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of services
	Items []Service `json:"items" protobuf:"bytes,2,rep,name=items"`
}
type Delete struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type User struct {
	username string `json:"username"`
	password string `json:"password"`
}
type Log struct {
	Time  string `orm:"pk;column(time)" json:"time"`
	Etype string `orm:"column(etype)" json:"etype"`
	Event string `orm:"column(event)" json:"event"`
}
type Response struct {
	Action string `json:"action"`
	Status string `json:"status"`
}
type PV struct {
	Name            string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	StorageName     string `json:"storageName"`
	AccessModes     string `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	StorageQuantity string `json:"storageQuantity"`
	NfsServer       string `json:"nfsserver"`
	NfsPath         string `json:"nfspath"`
}

type PersistentVolumeAccessMode string

const (
	// can be mounted read/write mode to exactly 1 host
	ReadWriteOnce PersistentVolumeAccessMode = "ReadWriteOnce"
	// can be mounted in read-only mode to many hosts
	ReadOnlyMany PersistentVolumeAccessMode = "ReadOnlyMany"
	// can be mounted in read/write mode to many hosts
	ReadWriteMany PersistentVolumeAccessMode = "ReadWriteMany"
)

type PVC struct {
	Name            string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	StorageName     string `json:"storageName"`
	AccessModes     string `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	StorageQuantity string `json:"storageQuantity"`
	NfsServer       string `json:"nfsserver"`
	NfsPath         string `json:"nfspath"`
	Namespace       string `json:"namespace"`
}
type Statefulset struct {
	Name        string      `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	StorageName string      `json:"storageName"`
	Namespace   string      `json:"namespace"`
	Replicas    int32       `json:"replicas"`
	Requests    string      `json:"requests"`
	Containers  []Container `json:"containers"`
	VMname      string      `json:"vmname"`
	MountPath   string      `json:"mountpath"`
	AccessModes string      `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
}
