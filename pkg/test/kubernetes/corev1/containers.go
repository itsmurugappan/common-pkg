package corev1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"knative.dev/pkg/ptr"
)

type expectedContainerOption func(*corev1.Container)

func ConstructExpectedContainerSpec(options ...expectedContainerOption) corev1.Container {
	cSpec := corev1.Container{}

	for _, fn := range options {
		fn(&cSpec)
	}

	return cSpec
}

func WithImage(image string) expectedContainerOption {
	return func(container *corev1.Container) {
		container.Image = image
	}
}

func WithEnv(keys []string, vals []string) expectedContainerOption {
	return func(container *corev1.Container) {
		var envs []corev1.EnvVar
		for i, key := range keys {
			envs = append(envs, corev1.EnvVar{
				Name:  key,
				Value: vals[i],
			})
		}
		container.Env = envs
	}
}

func WithEnvFromSecretorCM(names []string, types []string) expectedContainerOption {
	return func(container *corev1.Container) {
		container.EnvFrom = ConstructEnvFrom(names, types)
	}
}

func WithVolumeMounts(names []string, paths []string) expectedContainerOption {
	return func(container *corev1.Container) {
		container.VolumeMounts = ConstructMounts(names, paths)
	}
}

func WithPort(port int32) expectedContainerOption {
	return func(container *corev1.Container) {
		container.Ports = append(container.Ports, corev1.ContainerPort{
			ContainerPort: port,
		})
	}
}

func WithSecurityContext(user int64) expectedContainerOption {
	return func(container *corev1.Container) {
		container.SecurityContext = &corev1.SecurityContext{
			RunAsUser: ptr.Int64(user),
		}
	}
}

func WithName(name string) expectedContainerOption {
	return func(container *corev1.Container) {
		container.Name = name
	}
}

func WithCommand(cmd []string) expectedContainerOption {
	return func(container *corev1.Container) {
		if len(cmd) > 0 && cmd[0] != "" {
			container.Command = cmd
		}
	}
}

func WithImagePullPolicy(pullPolicy corev1.PullPolicy) expectedContainerOption {
	return func(container *corev1.Container) {
		if pullPolicy != "" {
			container.ImagePullPolicy = pullPolicy
		}
	}
}

func WithResources(cpuReq int64, cpuLim int64, memReq int64, memLim int64) expectedContainerOption {
	return func(container *corev1.Container) {
		resReq := corev1.ResourceRequirements{}
		if cpuReq > 0 || memReq > 0 {
			resReq.Requests = make(map[corev1.ResourceName]resource.Quantity)
		}
		if cpuLim > 0 || memLim > 0 {
			resReq.Limits = make(map[corev1.ResourceName]resource.Quantity)
		}
		if cpuReq != int64(0) {
			resReq.Requests[corev1.ResourceCPU] = *(resource.NewMilliQuantity(cpuReq, resource.DecimalSI))
		}
		if memReq != int64(0) {
			resReq.Requests[corev1.ResourceMemory] = *(resource.NewScaledQuantity(memReq, resource.Mega))
		}
		if cpuLim != int64(0) {
			resReq.Limits[corev1.ResourceCPU] = *(resource.NewMilliQuantity(cpuLim, resource.DecimalSI))
		}
		if memLim != int64(0) {
			resReq.Limits[corev1.ResourceMemory] = *(resource.NewScaledQuantity(memLim, resource.Mega))
		}
		container.Resources = resReq
	}
}
