package corev1

import (
	"gotest.tools/assert"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/itsmurugappan/kubernetes-resource-builder/pkg/kubernetes"
	teststubcorev1 "github.com/itsmurugappan/kubernetes-resource-builder/pkg/test/kubernetes/corev1"
)

func TestGetContainerSpec(t *testing.T) {
	for _, tc := range []struct {
		name          string
		wantContainer corev1.Container
		inputModel    kubernetes.ContainerSpec
		inputOptions  []containerSpecOption
	}{{
		name: "Container will All Options",
		wantContainer: teststubcorev1.ConstructExpectedContainerSpec(
			teststubcorev1.WithEnv([]string{"e1", "e2"}, []string{"v1", "v2"}),
			teststubcorev1.WithEnvFromSecretorCM([]string{"c1", "s1"}, []string{"CM", "Secret"}),
			teststubcorev1.WithVolumeMounts([]string{"c1", "s1"}, []string{"/p1", "/p2"}),
			teststubcorev1.WithPort(int32(8080)),
			teststubcorev1.WithPort(int32(9090)),
			teststubcorev1.WithSecurityContext(int64(1001)),
			teststubcorev1.WithName("foo"),
			teststubcorev1.WithImage("docker.com/bar"),
			teststubcorev1.WithCommand([]string{"python", "some.py"}),
			teststubcorev1.WithImagePullPolicy(corev1.PullAlways),
			teststubcorev1.WithResources(int64(10), int64(50), int64(128), int64(256))),
		inputModel: kubernetes.ContainerSpec{Image: "docker.com/bar"},
		inputOptions: []containerSpecOption{
			WithEnv([]corev1.EnvVar{corev1.EnvVar{Name: "e1", Value: "v1"}, corev1.EnvVar{Name: "e2", Value: "v2"}}),
			WithEnvFromSecretorCM([]kubernetes.EnvFrom{{"c1", "CM"}, {"s1", "Secret"}}),
			WithVolumeMounts(teststubcorev1.ConstructMounts([]string{"c1"}, []string{"/p1"}), teststubcorev1.ConstructMounts([]string{"s1"}, []string{"/p2"})),
			WithPort(int32(8080)),
			WithPort(int32(9090)),
			WithSecurityContext(int64(1001)),
			WithName("foo"),
			WithCommand([]string{"python", "some.py"}),
			WithImagePullPolicy(corev1.PullAlways),
			WithResources([]kubernetes.Resource{{"Requests", int64(10), int64(128)}, {"Limit", int64(50), int64(256)}}),
		},
	}, {
		name: "Container with null Options",
		wantContainer: teststubcorev1.ConstructExpectedContainerSpec(
			teststubcorev1.WithImage("docker.com/bar")),
		inputModel: kubernetes.ContainerSpec{Image: "docker.com/bar"},
		inputOptions: []containerSpecOption{
			WithEnv([]corev1.EnvVar{corev1.EnvVar{Name: "", Value: ""}}),
			WithEnvFromSecretorCM([]kubernetes.EnvFrom{{"", ""}}),
			WithVolumeMounts(teststubcorev1.ConstructMounts([]string{""}, []string{""}), teststubcorev1.ConstructMounts([]string{""}, []string{""})),
			WithPort(int32(0)),
			WithSecurityContext(int64(0)),
			WithName(""),
			WithCommand([]string{""}),
			WithResources([]kubernetes.Resource{{"", int64(0), int64(0)}}),
		},
	}, {
		name: "Container with specific resource options",
		wantContainer: teststubcorev1.ConstructExpectedContainerSpec(
			teststubcorev1.WithImage("docker.com/bar"),
			teststubcorev1.WithResources(int64(10), int64(0), int64(0), int64(256))),
		inputModel: kubernetes.ContainerSpec{Image: "docker.com/bar"},
		inputOptions: []containerSpecOption{
			WithResources([]kubernetes.Resource{{"Requests", int64(10), int64(0)}, {"Limit", int64(0), int64(256)}}),
		},
	}, {
		name: "Container with only resource requests",
		wantContainer: teststubcorev1.ConstructExpectedContainerSpec(
			teststubcorev1.WithImage("docker.com/bar"),
			teststubcorev1.WithResources(int64(10), int64(0), int64(128), int64(0))),
		inputModel: kubernetes.ContainerSpec{Image: "docker.com/bar"},
		inputOptions: []containerSpecOption{
			WithResources([]kubernetes.Resource{{"Requests", int64(10), int64(128)}, {"Limit", int64(0), int64(0)}}),
		},
	}, {
		name: "Container with only resource limits",
		wantContainer: teststubcorev1.ConstructExpectedContainerSpec(
			teststubcorev1.WithImage("docker.com/bar"),
			teststubcorev1.WithResources(int64(0), int64(10), int64(0), int64(128))),
		inputModel: kubernetes.ContainerSpec{Image: "docker.com/bar"},
		inputOptions: []containerSpecOption{
			WithResources([]kubernetes.Resource{{"Requests", int64(0), int64(0)}, {"Limit", int64(10), int64(128)}}),
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			actContainer := GetContainerSpec(tc.inputModel, tc.inputOptions...)
			assert.DeepEqual(t, &tc.wantContainer, &actContainer)
		})
	}
}
