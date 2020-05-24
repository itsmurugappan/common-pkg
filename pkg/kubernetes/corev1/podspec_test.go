package corev1

import (
	"gotest.tools/assert"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/itsmurugappan/kubernetes-resource-builder/pkg/kubernetes"
	teststubcorev1 "github.com/itsmurugappan/kubernetes-resource-builder/pkg/test/kubernetes/corev1"
)

func TestGetPodSpec(t *testing.T) {
	for _, tc := range []struct {
		name         string
		wantPodSpec  corev1.PodSpec
		inputModel   kubernetes.PodSpec
		inputOptions []PodSpecOption
	}{{
		name: "Pod will All Options",
		wantPodSpec: teststubcorev1.ConstructExpectedPodSpec(
			teststubcorev1.WithContainerOptions(
				teststubcorev1.WithEnv([]string{"e1", "e2"}, []string{"v1", "v2"}),
				teststubcorev1.WithEnvFromSecretorCM([]string{"c1", "s1"}, []string{"CM", "Secret"}),
				teststubcorev1.WithVolumeMounts([]string{"c1", "s1"}, []string{"/p1", "/p2"}),
				teststubcorev1.WithPort(int32(8080)),
				teststubcorev1.WithPort(int32(9090)),
				teststubcorev1.WithSecurityContext(int64(1001)),
				teststubcorev1.WithName("foo"),
				teststubcorev1.WithImage("docker.com/bar"),
				teststubcorev1.WithCommand([]string{"python", "some.py"}),
				teststubcorev1.WithResources(int64(10), int64(50), int64(128), int64(256))),
			teststubcorev1.WithVolumes([]string{"c1"}, []string{"s1"}),
			teststubcorev1.WithServiceAccount("admin-sa"),
			teststubcorev1.WithRestartPolicy("Never"),
		),
		inputModel: kubernetes.PodSpec{},
		inputOptions: []PodSpecOption{
			WithVolumes([]kubernetes.ContainerSpec{
				kubernetes.ContainerSpec{
					Secrets:    teststubcorev1.ConstructMounts([]string{"s1"}, []string{"/p2"}),
					ConfigMaps: teststubcorev1.ConstructMounts([]string{"c1"}, []string{"/p1"}),
				}}),
			WithServiceAccount("admin-sa"),
			WithRestartPolicy("Never"),
			WithContainerOptions(kubernetes.ContainerSpec{Image: "docker.com/bar"},
				WithEnv([]corev1.EnvVar{corev1.EnvVar{Name: "e1", Value: "v1"}, corev1.EnvVar{Name: "e2", Value: "v2"}}),
				WithEnvFromSecretorCM([]kubernetes.EnvFrom{{"c1", "CM"}, {"s1", "Secret"}}),
				WithVolumeMounts(teststubcorev1.ConstructMounts([]string{"c1"}, []string{"/p1"}), teststubcorev1.ConstructMounts([]string{"s1"}, []string{"/p2"})),
				WithPort(int32(8080)),
				WithPort(int32(9090)),
				WithSecurityContext(int64(1001)),
				WithName("foo"),
				WithCommand([]string{"python", "some.py"}),
				WithResources([]kubernetes.Resource{{"Requests", int64(10), int64(128)}, {"Limit", int64(50), int64(256)}})),
		},
	}, {
		name: "Pod with min options",
		wantPodSpec: teststubcorev1.ConstructExpectedPodSpec(
			teststubcorev1.WithContainerOptions(
				teststubcorev1.WithName("foo"),
				teststubcorev1.WithImage("docker.com/bar")),
		),
		inputModel: kubernetes.PodSpec{},
		inputOptions: []PodSpecOption{
			WithContainerOptions(kubernetes.ContainerSpec{Image: "docker.com/bar"},
				WithEnv([]corev1.EnvVar{corev1.EnvVar{Name: "", Value: ""}}),
				WithEnvFromSecretorCM([]kubernetes.EnvFrom{{"", ""}}),
				WithVolumeMounts(teststubcorev1.ConstructMounts([]string{""}, []string{""}), teststubcorev1.ConstructMounts([]string{""}, []string{""})),
				WithPort(int32(0)),
				WithSecurityContext(int64(0)),
				WithName("foo"),
				WithCommand([]string{""}),
				WithResources([]kubernetes.Resource{{"Requests", int64(0), int64(0)}, {"Limit", int64(0), int64(0)}})),
		},
	}, {
		name: "Pod with 2 containers",
		wantPodSpec: teststubcorev1.ConstructExpectedPodSpec(
			teststubcorev1.WithContainerOptions(
				teststubcorev1.WithName("foo"),
				teststubcorev1.WithPort(int32(8080)),
				teststubcorev1.WithVolumeMounts([]string{"c1", "s1"}, []string{"/p1", "/p2"}),
				teststubcorev1.WithImage("docker.com/foo")),
			teststubcorev1.WithContainerOptions(
				teststubcorev1.WithName("bar"),
				teststubcorev1.WithPort(int32(9090)),
				teststubcorev1.WithVolumeMounts([]string{"c2", "s2"}, []string{"/p1", "/p2"}),
				teststubcorev1.WithImage("docker.com/bar")),
			teststubcorev1.WithVolumes([]string{"c1"}, []string{"s1"}),
			teststubcorev1.WithVolumes([]string{"c2"}, []string{"s2"}),
		),
		inputModel: kubernetes.PodSpec{},
		inputOptions: []PodSpecOption{
			WithVolumes([]kubernetes.ContainerSpec{
				kubernetes.ContainerSpec{
					Secrets:    teststubcorev1.ConstructMounts([]string{"s1"}, []string{"/p2"}),
					ConfigMaps: teststubcorev1.ConstructMounts([]string{"c1"}, []string{"/p1"}),
				},
				kubernetes.ContainerSpec{
					Secrets:    teststubcorev1.ConstructMounts([]string{"s2"}, []string{"/p2"}),
					ConfigMaps: teststubcorev1.ConstructMounts([]string{"c2"}, []string{"/p1"}),
				}}),
			WithContainerOptions(kubernetes.ContainerSpec{Image: "docker.com/foo"},
				WithVolumeMounts(teststubcorev1.ConstructMounts([]string{"c1"}, []string{"/p1"}), teststubcorev1.ConstructMounts([]string{"s1"}, []string{"/p2"})),
				WithName("foo"),
				WithPort(int32(8080))),
			WithContainerOptions(kubernetes.ContainerSpec{Image: "docker.com/bar"},
				WithVolumeMounts(teststubcorev1.ConstructMounts([]string{"c2"}, []string{"/p1"}), teststubcorev1.ConstructMounts([]string{"s2"}, []string{"/p2"})),
				WithName("bar"),
				WithPort(int32(9090))),
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			actPod := GetPodSpec(tc.inputModel, tc.inputOptions...)
			assert.DeepEqual(t, &tc.wantPodSpec, &actPod)
		})
	}
}
