package batchv1

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/ptr"

	"github.com/itsmurugappan/kubernetes-resource-builder/test/kubernetes/corev1"
)

type expectedJobSpecOption func(*batchv1.Job)

func ConstructExpectedJobSpec(options ...expectedJobSpecOption) batchv1.Job {
	jobSpec := batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "foo",
		}}

	for _, fn := range options {
		fn(&jobSpec)
	}
	return jobSpec
}

func WithPodSpecOptions(options ...corev1.ExpectedPodSpecOption) expectedJobSpecOption {
	return func(job *batchv1.Job) {
		job.Spec.Template.Spec = corev1.ConstructExpectedPodSpec(options...)
	}
}

func WithTTL(ttl int32) expectedJobSpecOption {
	return func(job *batchv1.Job) {
		if ttl > int32(0) {
			job.Spec.TTLSecondsAfterFinished = ptr.Int32(ttl)
		}
	}
}

func WithBackoffLimit(backoffLimit int32) expectedJobSpecOption {
	return func(job *batchv1.Job) {
		if backoffLimit > int32(0) {
			job.Spec.BackoffLimit = ptr.Int32(backoffLimit)
		}
	}
}

func WithAnnotations(annotations map[string]string) expectedJobSpecOption {
	return func(job *batchv1.Job) {
		job.Spec.Template.ObjectMeta.Annotations = annotations
	}
}

func WithLabels(labels map[string]string) expectedJobSpecOption {
	return func(job *batchv1.Job) {
		job.Spec.Template.ObjectMeta.Labels = labels
	}
}

func WithActiveStatus() expectedJobSpecOption {
	return func(job *batchv1.Job) {
		job.Status.Active = int32(1)
	}
}
