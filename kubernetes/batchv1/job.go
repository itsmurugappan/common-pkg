package batchv1

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/ptr"

	"github.com/itsmurugappan/kubernetes-resource-builder/kubernetes"
	"github.com/itsmurugappan/kubernetes-resource-builder/kubernetes/corev1"
)

type jobSpecOption func(*batchv1.Job)

func GetJob(spec kubernetes.JobSpec, options ...jobSpecOption) batchv1.Job {
	jobSpec := batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: spec.Name,
		}}

	for _, fn := range options {
		fn(&jobSpec)
	}
	return jobSpec
}

func WithPodSpecOptions(podSpec kubernetes.PodSpec, options ...corev1.PodSpecOption) jobSpecOption {
	return func(job *batchv1.Job) {
		job.Spec.Template.Spec = corev1.GetPodSpec(podSpec, options...)
	}
}

func WithTTL(ttl int32) jobSpecOption {
	return func(job *batchv1.Job) {
		if ttl > int32(0) {
			job.Spec.TTLSecondsAfterFinished = ptr.Int32(ttl)
		}
	}
}

func WithBackoffLimit(backoffLimit int32) jobSpecOption {
	return func(job *batchv1.Job) {
		if backoffLimit > int32(0) {
			job.Spec.BackoffLimit = ptr.Int32(backoffLimit)
		}
	}
}

func WithAnnotations(inAnnotations []kubernetes.KV) jobSpecOption {
	return func(job *batchv1.Job) {
		if len(inAnnotations) > 0 && inAnnotations[0].Key != "" {
			annotations := make(map[string]string)
			for _, ann := range inAnnotations {
				annotations[ann.Key] = ann.Value
			}
			job.Spec.Template.ObjectMeta.Annotations = annotations
		}
	}
}

func WithLabels(inLabels []kubernetes.KV) jobSpecOption {
	return func(job *batchv1.Job) {
		if len(inLabels) > 0 && inLabels[0].Key != "" {
			labels := make(map[string]string)
			for _, label := range inLabels {
				labels[label.Key] = label.Value
			}
			job.Spec.Template.ObjectMeta.Labels = labels
		}
	}
}
