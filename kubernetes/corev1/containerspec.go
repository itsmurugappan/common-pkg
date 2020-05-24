package corev1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"knative.dev/pkg/ptr"

	"github.com/itsmurugappan/kubernetes-resource-builder/kubernetes"
)

type containerSpecOption func(*corev1.Container)

func GetContainerSpec(spec kubernetes.ContainerSpec, options ...containerSpecOption) corev1.Container {

	cSpec := corev1.Container{
		Image: spec.Image,
	}
	for _, fn := range options {
		fn(&cSpec)
	}
	return cSpec
}

func WithEnv(envs []corev1.EnvVar) containerSpecOption {
	return func(container *corev1.Container) {
		if len(envs) > 0 && envs[0].Name != "" {
			container.Env = append(container.Env, envs...)
		}
	}
}

func WithEnvFromSecretorCM(envFromSecretorCM []kubernetes.EnvFrom) containerSpecOption {
	return func(container *corev1.Container) {
		envList := GetEnvfromSecretorCM(envFromSecretorCM)
		container.EnvFrom = envList
	}
}

func WithVolumeMounts(cms []corev1.VolumeMount, secrets []corev1.VolumeMount) containerSpecOption {
	return func(container *corev1.Container) {
		mountList := GetVolumeMounts(cms, secrets)
		container.VolumeMounts = mountList
	}
}

func WithPort(port int32) containerSpecOption {
	return func(container *corev1.Container) {
		if port > 0 {
			container.Ports = append(container.Ports, corev1.ContainerPort{
				ContainerPort: port,
			})
		}
	}
}

func WithSecurityContext(user int64) containerSpecOption {
	return func(container *corev1.Container) {
		if user > 0 {
			container.SecurityContext = &corev1.SecurityContext{
				RunAsUser: ptr.Int64(user),
			}
		}
	}
}

func WithName(name string) containerSpecOption {
	return func(container *corev1.Container) {
		if name != "" {
			container.Name = name
		}
	}
}

func WithCommand(cmd []string) containerSpecOption {
	return func(container *corev1.Container) {
		if len(cmd) > 0 && cmd[0] != "" {
			container.Command = cmd
		}
	}
}

func WithImagePullPolicy(pullPolicy corev1.PullPolicy) containerSpecOption {
	return func(container *corev1.Container) {
		if pullPolicy != "" {
			container.ImagePullPolicy = pullPolicy
		}
	}
}

func WithResources(resources []kubernetes.Resource) containerSpecOption {
	return func(container *corev1.Container) {
		if len(resources) > 0 {
			resReq := corev1.ResourceRequirements{}
			for _, res := range resources {
				switch res.Type {
				case "Requests":
					if res.CPU != int64(0) || res.Mem != int64(0) {
						resReq.Requests = make(map[corev1.ResourceName]resource.Quantity)
					}
					if res.CPU != int64(0) {
						resReq.Requests[corev1.ResourceCPU] = *(resource.NewMilliQuantity(res.CPU, resource.DecimalSI))
					}
					if res.Mem != int64(0) {
						resReq.Requests[corev1.ResourceMemory] = *(resource.NewScaledQuantity(res.Mem, resource.Mega))
					}
				case "Limit":
					if res.CPU != int64(0) || res.Mem != int64(0) {
						resReq.Limits = make(map[corev1.ResourceName]resource.Quantity)
					}
					if res.CPU != int64(0) {
						resReq.Limits[corev1.ResourceCPU] = *(resource.NewMilliQuantity(res.CPU, resource.DecimalSI))
					}
					if res.Mem != int64(0) {
						resReq.Limits[corev1.ResourceMemory] = *(resource.NewScaledQuantity(res.Mem, resource.Mega))
					}
				}
			}
			container.Resources = resReq
		}
	}
}
