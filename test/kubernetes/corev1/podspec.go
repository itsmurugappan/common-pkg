package corev1

import (
	corev1 "k8s.io/api/core/v1"
)

type ExpectedPodSpecOption func(*corev1.PodSpec)

func ConstructExpectedPodSpec(options ...ExpectedPodSpecOption) corev1.PodSpec {
	pSpec := corev1.PodSpec{}

	for _, fn := range options {
		fn(&pSpec)
	}

	return pSpec
}

func WithContainerOptions(options ...expectedContainerOption) ExpectedPodSpecOption {
	return func(podSpec *corev1.PodSpec) {
		podSpec.Containers = append(podSpec.Containers, ConstructExpectedContainerSpec(options...))
	}
}

func WithVolumes(cms []string, secrets []string) ExpectedPodSpecOption {
	return func(spec *corev1.PodSpec) {
		spec.Volumes = append(spec.Volumes, ConstructVolumeSources(cms, secrets)...)
	}
}

func WithServiceAccount(sa string) ExpectedPodSpecOption {
	return func(spec *corev1.PodSpec) {
		spec.ServiceAccountName = sa
	}
}

func WithRestartPolicy(policy string) ExpectedPodSpecOption {
	return func(spec *corev1.PodSpec) {
		spec.RestartPolicy = corev1.RestartPolicy(policy)
	}
}
