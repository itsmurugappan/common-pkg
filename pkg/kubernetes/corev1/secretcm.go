package corev1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	types "github.com/itsmurugappan/kubernetes-resource-builder/pkg/kubernetes"
)

const (
	//CM_MISSING - error message to indicate CM is missing
	CM_MISSING = "cm %s is not found in the namespace %s, create cm before referring in the function"
	//SECRET_MISSING - error message to indicate Secret is missing
	SECRET_MISSING = "secret %s is not found in the namespace %s, create secret before referring in the function"
	//INVALID_ENV_FROM_TYPE - error message to indicate invalid type to mount as env
	INVALID_ENV_FROM_TYPE = "Provide a valid EnvFrom type. Should be 'CM' or 'Secret'"
)

func CheckIfCMExist(nsName string, cm string, typedcorev1 typedcorev1.CoreV1Interface) bool {
	getOpts := metav1.GetOptions{TypeMeta: metav1.TypeMeta{Kind: "ConfigMap", APIVersion: "v1"}}

	if _, err := typedcorev1.ConfigMaps(nsName).Get(cm, getOpts); err != nil {
		return false
	}

	return true
}

func CheckIfSecretExist(nsName string, secret string, typedcorev1 typedcorev1.CoreV1Interface) bool {
	getOpts := metav1.GetOptions{TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"}}

	if _, err := typedcorev1.Secrets(nsName).Get(secret, getOpts); err != nil {
		return false
	}

	return true
}

func CheckSecretMounts(ns string, typedcorev1 typedcorev1.CoreV1Interface, secrets []corev1.VolumeMount) error {
	// check secret
	for _, secret := range secrets {
		if secret.Name == "" {
			continue
		}
		if !CheckIfSecretExist(ns, secret.Name, typedcorev1) {
			return fmt.Errorf(SECRET_MISSING, secret.Name, ns)
		}
	}
	return nil
}

func CheckCMMounts(ns string, typedcorev1 typedcorev1.CoreV1Interface, cms []corev1.VolumeMount) error {
	// check cm
	for _, cm := range cms {
		if cm.Name == "" {
			continue
		}
		if !CheckIfCMExist(ns, cm.Name, typedcorev1) {
			return fmt.Errorf(CM_MISSING, cm.Name, ns)
		}
	}
	return nil
}

func CheckEnvFromResources(ns string, typedcorev1 typedcorev1.CoreV1Interface, envsFrom []types.EnvFrom) error {
	// check env from
	for _, env := range envsFrom {
		if env.Name == "" {
			continue
		}
		switch env.Type {
		case "CM":
			if !CheckIfCMExist(ns, env.Name, typedcorev1) {
				return fmt.Errorf(CM_MISSING, env.Name, ns)
			}
		case "Secret":
			if !CheckIfSecretExist(ns, env.Name, typedcorev1) {
				return fmt.Errorf(SECRET_MISSING, env.Name, ns)
			}
		default:
			return fmt.Errorf(INVALID_ENV_FROM_TYPE)
		}
	}
	return nil
}

func GetSecrets(nsName, secret string, typedcorev1 typedcorev1.CoreV1Interface) (*corev1.Secret, error) {
	getOpts := metav1.GetOptions{TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"}}

	return typedcorev1.Secrets(nsName).Get(secret, getOpts)
}

func GetSAToken(nsName, saName string, typedcorev1 typedcorev1.CoreV1Interface) (string, error) {
	listOpts := metav1.ListOptions{FieldSelector: "type=kubernetes.io/service-account-token"}

	secretList, err := typedcorev1.Secrets(nsName).List(listOpts)
	if err != nil {
		return "", err
	}
	for _, secret := range secretList.Items {
		if secret.ObjectMeta.Annotations["kubernetes.io/service-account.name"] == saName {
			return string(secret.Data["token"]), nil
		}
	}
	return "", fmt.Errorf("No secret for sa %s", saName)
}
