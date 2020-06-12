package corev1

import (
	"fmt"
	"testing"

	"gotest.tools/assert"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	testclient "k8s.io/client-go/kubernetes/fake"

	types "github.com/itsmurugappan/kubernetes-resource-builder/pkg/kubernetes"
	teststubcorev1 "github.com/itsmurugappan/kubernetes-resource-builder/pkg/test/kubernetes/corev1"
)

func TestCheckIfCMExist(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           bool
		input          []string
		runtimeObjects []runtime.Object
	}{{
		name:           "cm created",
		want:           true,
		input:          []string{"foo", "cm1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructConfigMap("foo", "cm1")},
	}, {
		name:           "cm not created",
		want:           false,
		input:          []string{"foo", "cm1"},
		runtimeObjects: []runtime.Object{},
	}, {
		name:           "name mismatch",
		want:           false,
		input:          []string{"foo", "cm1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructConfigMap("foo", "cm2")},
	}, {
		name:           "namespace mismatch",
		want:           false,
		input:          []string{"foo", "cm1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructConfigMap("baz", "cm1")},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			act := CheckIfCMExist(tc.input[0], tc.input[1], testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1())
			assert.DeepEqual(t, &tc.want, &act)
		})
	}
}

func TestCheckIfSecretExist(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           bool
		input          []string
		runtimeObjects []runtime.Object
	}{{
		name:           "secret created",
		want:           true,
		input:          []string{"foo", "s1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructSecret("foo", "s1")},
	}, {
		name:           "secret not created",
		want:           false,
		input:          []string{"foo", "s1"},
		runtimeObjects: []runtime.Object{},
	}, {
		name:           "name mismatch",
		want:           false,
		input:          []string{"foo", "s1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructSecret("foo", "s2")},
	}, {
		name:           "namespace mismatch",
		want:           false,
		input:          []string{"foo", "s1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructSecret("baz", "s1")},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			act := CheckIfSecretExist(tc.input[0], tc.input[1], testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1())
			assert.DeepEqual(t, &tc.want, &act)
		})
	}
}

func TestCheckCMMounts(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           string
		ns             string
		input          []corev1.VolumeMount
		runtimeObjects []runtime.Object
	}{{
		name:  "2 cm created",
		want:  "",
		ns:    "foo",
		input: teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "cm1"),
			teststubcorev1.ConstructConfigMap("foo", "cm2"),
		},
	}, {
		name:  "1 cm not created",
		want:  fmt.Sprintf(CM_MISSING, "cm1", "foo"),
		ns:    "foo",
		input: teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "cm2"),
		},
	}, {
		name:  "name mismatch",
		want:  fmt.Sprintf(CM_MISSING, "cm3", "foo"),
		ns:    "foo",
		input: teststubcorev1.ConstructMounts([]string{"cm3", "cm4"}, []string{"p1", "p2"}),
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "cm1"),
			teststubcorev1.ConstructConfigMap("foo", "cm2"),
		},
	}, {
		name:  "namespace mismatch",
		want:  fmt.Sprintf(CM_MISSING, "cm1", "default"),
		input: teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		ns:    "default",
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "cm1"),
			teststubcorev1.ConstructConfigMap("foo", "cm2"),
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckCMMounts(tc.ns, testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1(), tc.input)
			if tc.want == "" {
				assert.NilError(t, err)
			} else {
				assert.Error(t, err, tc.want)
			}
		})
	}
}

func TestCheckSecretMounts(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           string
		ns             string
		input          []corev1.VolumeMount
		runtimeObjects []runtime.Object
	}{{
		name:  "2 secrets created",
		want:  "",
		ns:    "foo",
		input: teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p1", "p2"}),
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructSecret("foo", "s1"),
			teststubcorev1.ConstructSecret("foo", "s2"),
		},
	}, {
		name:  "1 cm not created",
		want:  fmt.Sprintf(SECRET_MISSING, "s1", "foo"),
		ns:    "foo",
		input: teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p1", "p2"}),
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructSecret("foo", "s2"),
		},
	}, {
		name:  "name mismatch",
		want:  fmt.Sprintf(SECRET_MISSING, "s3", "foo"),
		ns:    "foo",
		input: teststubcorev1.ConstructMounts([]string{"s3", "s4"}, []string{"p1", "p2"}),
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructSecret("foo", "s1"),
			teststubcorev1.ConstructSecret("foo", "s2"),
		},
	}, {
		name:  "namespace mismatch",
		want:  fmt.Sprintf(SECRET_MISSING, "s1", "default"),
		input: teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p1", "p2"}),
		ns:    "default",
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructSecret("foo", "s1"),
			teststubcorev1.ConstructSecret("foo", "s2"),
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckSecretMounts(tc.ns, testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1(), tc.input)
			if tc.want == "" {
				assert.NilError(t, err)
			} else {
				assert.Error(t, err, tc.want)
			}
		})
	}
}

func TestEnvFromResources(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           string
		ns             string
		input          []types.EnvFrom
		runtimeObjects []runtime.Object
	}{{
		name:  "1 cm 2 secret",
		want:  "",
		ns:    "foo",
		input: []types.EnvFrom{{"c1", "CM"}, {"s1", "Secret"}, {"s2", "Secret"}},
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "c1"),
			teststubcorev1.ConstructSecret("foo", "s1"),
			teststubcorev1.ConstructSecret("foo", "s2"),
		},
	}, {
		name:  "cm missing",
		want:  fmt.Sprintf(CM_MISSING, "c1", "foo"),
		ns:    "foo",
		input: []types.EnvFrom{{"s1", "Secret"}, {"c1", "CM"}},
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructSecret("foo", "s1"),
		},
	}, {
		name:  "secret missing",
		want:  fmt.Sprintf(SECRET_MISSING, "s1", "foo"),
		ns:    "foo",
		input: []types.EnvFrom{{"cm1", "CM"}, {"s1", "Secret"}},
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "cm1"),
		},
	}, {
		name:  "wrong type",
		want:  fmt.Sprintf(INVALID_ENV_FROM_TYPE),
		input: []types.EnvFrom{{"c1", "CM"}, {"s2", "Secret"}, {"cm2", "config"}},
		ns:    "foo",
		runtimeObjects: []runtime.Object{
			teststubcorev1.ConstructConfigMap("foo", "c1"),
			teststubcorev1.ConstructSecret("foo", "s2"),
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckEnvFromResources(tc.ns, testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1(), tc.input)
			if tc.want == "" {
				assert.NilError(t, err)
			} else {
				assert.Error(t, err, tc.want)
			}
		})
	}
}

func TestGetSecrets(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           *corev1.Secret
		input          []string
		runtimeObjects []runtime.Object
	}{{
		name:           "secret created",
		want:           teststubcorev1.ConstructSecret("foo", "s1"),
		input:          []string{"foo", "s1"},
		runtimeObjects: []runtime.Object{teststubcorev1.ConstructSecret("foo", "s1")},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			secret, err := GetSecrets(tc.input[0], tc.input[1], testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1())
			assert.NilError(t, err)
			assert.DeepEqual(t, tc.want, secret)
		})
	}
}

func TestGetSAToken(t *testing.T) {
	for _, tc := range []struct {
		name           string
		want           string
		input          []string
		runtimeObjects []runtime.Object
		isErr          bool
	}{{
		name:  "happy path",
		want:  "1asdadasd1",
		input: []string{"foo", "sa1"},
		runtimeObjects: []runtime.Object{teststubcorev1.GetSecret("foo", "sa1",
			teststubcorev1.WithSecretType("kubernetes.io/service-account-token"),
			teststubcorev1.WithSecretAnnotations(map[string]string{"kubernetes.io/service-account.name": "sa1"}),
		)},
	}, {
		name:  "no sa",
		want:  "No secret for sa sa2",
		isErr: true,
		input: []string{"foo", "sa2"},
		runtimeObjects: []runtime.Object{teststubcorev1.GetSecret("foo", "sa1",
			teststubcorev1.WithSecretType("kubernetes.io/service-account-token"),
			teststubcorev1.WithSecretAnnotations(map[string]string{"kubernetes.io/service-account.name": "sa1"}),
		)},
	}, {
		name:           "no secrets",
		want:           "No secret for sa sa2",
		isErr:          true,
		input:          []string{"foo", "sa2"},
		runtimeObjects: []runtime.Object{},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			secret, err := GetSAToken(tc.input[0], tc.input[1], testclient.NewSimpleClientset((tc.runtimeObjects)...).CoreV1())
			if tc.isErr {
				assert.Error(t, err, tc.want)
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tc.want, secret)
			}
		})
	}
}
