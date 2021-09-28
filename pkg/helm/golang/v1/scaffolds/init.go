package scaffolds

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/golang/v1/scaffolds/templates"
)

const (
	EnvtestK8sVersion        = "1.21"
	ControllerToolsVersion   = "v0.6.1"
	ControllerRuntimeVersion = "v0.9.2"
	HelmVersion              = "v3.7.0"
)

type InitScaffolder struct {
	Config config.Config
	fs     machinery.Filesystem
}

func (s *InitScaffolder) InjectFS(fs machinery.Filesystem) {
	s.fs = fs
}

func (s *InitScaffolder) Scaffold() error {
	scaffold := machinery.NewScaffold(s.fs, machinery.WithConfig(s.Config))
	return scaffold.Execute(
		&templates.Makefile{
			EnvtestK8sVersion:        EnvtestK8sVersion,
			ControllerToolsVersion:   ControllerToolsVersion,
			ControllerRuntimeVersion: ControllerRuntimeVersion,
			HelmVersion:              HelmVersion,
		},
	)
}
