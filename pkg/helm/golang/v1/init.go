package v1

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/golang/v1/scaffolds"
)

type initSubcommand struct {
	config config.Config
}

func (s *initSubcommand) InjectConfig(c config.Config) error {
	s.config = c
	return nil
}

func (s *initSubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := &scaffolds.InitScaffolder{Config: s.config}
	scaffolder.InjectFS(fs)
	return scaffolder.Scaffold()
}
