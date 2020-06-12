package kubernetes

import (
	"fmt"
	"io/ioutil"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//GetInClusterKubeClient get kube context based on the service account
//on the deplioyment
func GetInClusterKubeClient() (*kubernetes.Clientset, error) {
	config, cfgErr := rest.InClusterConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return kubernetes.NewForConfig(config)
}

//GetCurrentNamespace return the namespace where pod is running
func GetCurrentNamespace() string {
	dat, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		fmt.Println(err)
	}

	return string(dat)
}
