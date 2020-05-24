package corev1

import (
	"gotest.tools/assert"
	"testing"

	corev1 "k8s.io/api/core/v1"

	teststubcorev1 "github.com/itsmurugappan/kubernetes-resource-builder/test/kubernetes/corev1"
)

func TestGetVolumeSourcesandMounts(t *testing.T) {
	for _, tc := range []struct {
		name            string
		wantVolume      []corev1.Volume
		wantVolumeMount []corev1.VolumeMount
		inputCM         []corev1.VolumeMount
		inputSecret     []corev1.VolumeMount
	}{{
		name:            "happy path - cm and secret volume",
		wantVolume:      teststubcorev1.ConstructMergedVols([]string{"cm1", "cm2"}, []string{"s1", "s2"}),
		wantVolumeMount: teststubcorev1.ConstructMounts([]string{"cm1", "cm2", "s1", "s2"}, []string{"p1", "p2", "p3", "p4"}),
		inputCM:         teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		inputSecret:     teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p3", "p4"}),
	}, {
		name:            "happy path - cm and secret volume without validation",
		wantVolume:      teststubcorev1.ConstructMergedVols([]string{"cm1", "cm2"}, []string{"s1", "s2"}),
		wantVolumeMount: teststubcorev1.ConstructMounts([]string{"cm1", "cm2", "s1", "s2"}, []string{"p1", "p2", "p3", "p4"}),
		inputCM:         teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		inputSecret:     teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p3", "p4"}),
	}, {
		name:            "cm and secret volume with error",
		wantVolume:      teststubcorev1.ConstructMergedVols([]string{"cm1", "cm2"}, []string{"s1", "s2"}),
		wantVolumeMount: teststubcorev1.ConstructMounts([]string{"cm1", "cm2", "s1", "s2"}, []string{"p1", "p2", "p3", "p4"}),
		inputCM:         teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		inputSecret:     teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p3", "p4"}),
	}, {
		name:            "cm with error",
		wantVolume:      teststubcorev1.ConstructMergedVols([]string{"cm1", "cm2"}, []string{}),
		wantVolumeMount: teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		inputCM:         teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
	}, {
		name:            "empty cm and secret",
		wantVolume:      teststubcorev1.ConstructMergedVols(nil, nil),
		wantVolumeMount: teststubcorev1.ConstructMounts(nil, nil),
		inputCM:         teststubcorev1.ConstructMounts([]string{""}, []string{""}),
		inputSecret:     teststubcorev1.ConstructMounts([]string{""}, []string{""}),
	}, {
		name:            "only cm",
		wantVolume:      teststubcorev1.ConstructMergedVols([]string{"cm1", "cm2"}, nil),
		wantVolumeMount: teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
		inputCM:         teststubcorev1.ConstructMounts([]string{"cm1", "cm2"}, []string{"p1", "p2"}),
	}, {
		name:            "only secret",
		wantVolume:      teststubcorev1.ConstructMergedVols(nil, []string{"s1", "s2"}),
		wantVolumeMount: teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p1", "p2"}),
		inputSecret:     teststubcorev1.ConstructMounts([]string{"s1", "s2"}, []string{"p1", "p2"}),
	}} {
		t.Run(tc.name, func(t *testing.T) {
			actVol := GetVolumeSources(tc.inputCM, tc.inputSecret)
			assert.DeepEqual(t, &tc.wantVolume, &actVol)
			actMt := GetVolumeMounts(tc.inputCM, tc.inputSecret)
			assert.DeepEqual(t, &tc.wantVolumeMount, &actMt)
		})
	}
}
