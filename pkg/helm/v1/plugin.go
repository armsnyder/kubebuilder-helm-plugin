package v1

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/model/stage"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm"
	basev1 "github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/base/v1"
	golangv1 "github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/golang/v1"
)

func New() plugin.Bundle {
	bundle, err := plugin.NewBundle(helm.Domain, plugin.Version{Number: 1, Stage: stage.Alpha},
		&basev1.Plugin{},
		&golangv1.Plugin{},
	)
	if err != nil {
		panic(err)
	}
	return bundle
}
