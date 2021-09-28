package scaffolds

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/base/v1/scaffolds/templates/chart"
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
		&chart.Chart{ProjectName: s.Config.GetProjectName()},
		&chart.Values{ProjectName: s.Config.GetProjectName()},
		&chart.HelmIgnore{ProjectName: s.Config.GetProjectName()},
	)
}
