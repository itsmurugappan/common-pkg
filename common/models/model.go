package models

import (
	corev1 "k8s.io/api/core/v1"
)

type EnvFrom struct {
	Name string
	Type string
}

type ContainerSpec struct {
	Image             string
	Port              int32
	Name              string
	Resources         []Resource
	Secrets           []corev1.VolumeMount
	ConfigMaps        []corev1.VolumeMount
	EnvVariables      []corev1.EnvVar
	User              int64
	EnvFromSecretorCM []EnvFrom
	Cmd               []string
}

type PodSpec struct {
	Containers []ContainerSpec
}

type JobSpec struct {
	Spec PodSpec
	Name string
}

type Resource struct {
	Type string
	CPU  int64
	Mem  int64
}

type KV struct {
	Key   string
	Value string
}
