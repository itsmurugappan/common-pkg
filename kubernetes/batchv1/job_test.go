package batchv1

import (
	"gotest.tools/assert"
	"testing"

	batchv1 "k8s.io/api/batch/v1"

	"github.com/itsmurugappan/common-pkg/common/models"
	corev1 "github.com/itsmurugappan/common-pkg/kubernetes/corev1"
	teststubbatchv1 "github.com/itsmurugappan/common-pkg/test/kubernetes/batchv1"
	teststubcorev1 "github.com/itsmurugappan/common-pkg/test/kubernetes/corev1"
)

func TestGetPodSpec(t *testing.T) {
	for _, tc := range []struct {
		name         string
		wantJob      batchv1.Job
		inputModel   models.JobSpec
		inputOptions []jobSpecOption
	}{{
		name: "Job with All Options",
		wantJob: teststubbatchv1.ConstructExpectedJobSpec(
			teststubbatchv1.WithPodSpecOptions(
				teststubcorev1.WithServiceAccount("admin-sa"),
				teststubcorev1.WithRestartPolicy("Never")),
			teststubbatchv1.WithTTL(int32(100)),
			teststubbatchv1.WithAnnotations(map[string]string{"key1": "val1", "key2": "val2"}),
			teststubbatchv1.WithLabels(map[string]string{"key1": "val1", "key2": "val2"}),
			teststubbatchv1.WithBackoffLimit(int32(1)),
		),
		inputModel: models.JobSpec{Name: "foo"},
		inputOptions: []jobSpecOption{
			WithTTL(int32(100)),
			WithAnnotations([]models.KV{{"key1", "val1"}, {"key2", "val2"}}),
			WithLabels([]models.KV{{"key1", "val1"}, {"key2", "val2"}}),
			WithBackoffLimit(int32(1)),
			WithPodSpecOptions(models.PodSpec{},
				corev1.WithServiceAccount("admin-sa"),
				corev1.WithRestartPolicy("Never")),
		},
	}, {
		name: "Job with null options",
		wantJob: teststubbatchv1.ConstructExpectedJobSpec(
			teststubbatchv1.WithPodSpecOptions(
				teststubcorev1.WithServiceAccount("admin-sa"),
				teststubcorev1.WithRestartPolicy("Never")),
		),
		inputModel: models.JobSpec{Name: "foo"},
		inputOptions: []jobSpecOption{
			WithTTL(int32(0)),
			WithBackoffLimit(int32(0)),
			WithPodSpecOptions(models.PodSpec{},
				corev1.WithServiceAccount("admin-sa"),
				corev1.WithRestartPolicy("Never")),
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			actJob := GetJob(tc.inputModel, tc.inputOptions...)
			assert.DeepEqual(t, &tc.wantJob, &actJob)
		})
	}
}
