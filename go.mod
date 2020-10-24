module github.com/itsmurugappan/kubernetes-resource-builder

go 1.15

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.17.4
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
	knative.dev/pkg v0.0.0-20200417160248-9320e44d1bf7
)

replace (
	k8s.io/api => k8s.io/api v0.18.8

	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
)
