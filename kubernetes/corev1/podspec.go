package corev1

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/itsmurugappan/common-pkg/common/models"
)

type PodSpecOption func(*corev1.PodSpec)

func GetPodSpec(spec models.PodSpec, options ...PodSpecOption) corev1.PodSpec {
	podSpec := corev1.PodSpec{}

	for _, fn := range options {
		fn(&podSpec)
	}
	return podSpec
}

func WithContainerOptions(cspec models.ContainerSpec, options ...containerSpecOption) PodSpecOption {
	return func(podSpec *corev1.PodSpec) {
		podSpec.Containers = append(podSpec.Containers, GetContainerSpec(cspec, options...))
	}
}

func WithVolumes(containers []models.ContainerSpec) PodSpecOption {
	return func(spec *corev1.PodSpec) {
		for _, container := range containers {
			volList := GetVolumeSources(container.ConfigMaps, container.Secrets)
			if len(volList) > 0 {
				spec.Volumes = append(spec.Volumes, volList...)
			}
		}
	}
}

func WithServiceAccount(sa string) PodSpecOption {
	return func(spec *corev1.PodSpec) {
		if sa != "" {
			spec.ServiceAccountName = sa
		}
	}
}

func WithRestartPolicy(policy string) PodSpecOption {
	return func(spec *corev1.PodSpec) {
		if policy != "" {
			spec.RestartPolicy = corev1.RestartPolicy(policy)
		}
	}
}
