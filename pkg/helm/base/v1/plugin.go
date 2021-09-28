package v1

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	configv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v3/pkg/model/stage"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm"
)

type Plugin struct {
	initSubcommand
}

func (Plugin) Name() string {
	return helm.Domain
}

func (Plugin) Version() plugin.Version {
	return plugin.Version{Number: 1, Stage: stage.Alpha}
}

func (Plugin) SupportedProjectVersions() []config.Version {
	return []config.Version{configv3.Version}
}

func (p *Plugin) GetInitSubcommand() plugin.InitSubcommand {
	return &p.initSubcommand
}
