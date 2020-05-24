package corev1

import (
	"gotest.tools/assert"
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/itsmurugappan/kubernetes-resource-builder/pkg/kubernetes"
	teststubcorev1 "github.com/itsmurugappan/kubernetes-resource-builder/pkg/test/kubernetes/corev1"
)

func TestEnvFromSecretorCM(t *testing.T) {
	for _, tc := range []struct {
		name  string
		want  []corev1.EnvFromSource
		input []kubernetes.EnvFrom
	}{{
		name:  "test cm only - happy path",
		want:  teststubcorev1.ConstructEnvFrom([]string{"cm1"}, []string{"CM"}),
		input: []kubernetes.EnvFrom{{"cm1", "CM"}},
	}, {
		name:  "test multiple cm without creating cm and expect error",
		want:  teststubcorev1.ConstructEnvFrom([]string{"cm1", "cm2"}, []string{"CM", "CM"}),
		input: []kubernetes.EnvFrom{{"cm1", "CM"}, {"cm2", "CM"}},
	}, {
		name:  "test multiple secret w/o creating a secret and expect err",
		want:  teststubcorev1.ConstructEnvFrom([]string{"s1", "s2"}, []string{"Secret", "Secret"}),
		input: []kubernetes.EnvFrom{{"s1", "Secret"}, {"s2", "Secret"}},
	}, {
		name:  "test multiple secret and cm without creating cm & secret - no error",
		want:  teststubcorev1.ConstructEnvFrom([]string{"s1", "s2", "cm1", "cm2"}, []string{"Secret", "Secret", "CM", "CM"}),
		input: []kubernetes.EnvFrom{{"s1", "Secret"}, {"s2", "Secret"}, {"cm1", "CM"}, {"cm2", "CM"}},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			actEnvFrom := GetEnvfromSecretorCM(tc.input)
			assert.DeepEqual(t, &tc.want, &actEnvFrom)
		})
	}
}

func TestGetEnvFromHTTPParam(t *testing.T) {
	for _, tc := range []struct {
		name  string
		want  []corev1.EnvVar
		input map[string][]string
	}{{
		name:  "test 2 query params",
		want:  []corev1.EnvVar{corev1.EnvVar{Name: "e1", Value: "v1;v2"}},
		input: map[string][]string{"e1": []string{"v1", "v2"}},
	}, {
		name:  "test empty",
		want:  nil,
		input: make(map[string][]string),
	}} {
		t.Run(tc.name, func(t *testing.T) {
			actEnv := GetEnvFromHTTPParam(tc.input)
			assert.DeepEqual(t, &tc.want, &actEnv)
		})
	}
}
