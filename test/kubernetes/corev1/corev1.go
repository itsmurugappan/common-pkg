package corev1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ConstructEnvFrom(names []string, types []string) []corev1.EnvFromSource {
	var envs []corev1.EnvFromSource
	for i, name := range names {
		if types[i] == "Secret" {
			envs = append(envs, corev1.EnvFromSource{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
				},
			})
		} else {
			envs = append(envs, corev1.EnvFromSource{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name,
					},
				},
			})
		}
	}
	return envs
}

func ConstructSecret(ns string, name string) *corev1.Secret {
	return &corev1.Secret{
		TypeMeta:   metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		Data:       map[string][]byte{"token": []byte("1asdadasd1")},
	}
}

func ConstructConfigMap(ns string, name string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		TypeMeta:   metav1.TypeMeta{Kind: "ConfigMap", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
	}
}

func ConstructMounts(names []string, paths []string) []corev1.VolumeMount {
	var mounts []corev1.VolumeMount
	for i, name := range names {
		mounts = append(mounts, corev1.VolumeMount{
			Name:      name,
			MountPath: paths[i],
		})
	}
	return mounts
}

func ConstructSecretVols(volNames []string) []corev1.Volume {
	var vols []corev1.Volume
	for _, volName := range volNames {
		vols = append(vols, corev1.Volume{
			Name: volName, VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: volName,
				},
			},
		})
	}
	return vols
}

func ConstructCMVols(volNames []string) []corev1.Volume {
	var vols []corev1.Volume
	for _, volName := range volNames {
		vols = append(vols, corev1.Volume{
			Name: volName, VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: volName,
					},
				},
			},
		})
	}
	return vols
}

func ConstructVolumeSources(cms []string, secrets []string) []corev1.Volume {
	var volList []corev1.Volume
	volList = append(volList, ConstructCMVols(cms)...)
	volList = append(volList, ConstructSecretVols(secrets)...)
	return volList
}

func ConstructMergedVols(cmVolNames []string, secretVolNames []string) []corev1.Volume {
	var vols []corev1.Volume
	vols = append(vols, ConstructCMVols(cmVolNames)...)
	vols = append(vols, ConstructSecretVols(secretVolNames)...)
	return vols
}
