package main

import (
	"log"

	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
	configv3 "sigs.k8s.io/kubebuilder/v3/pkg/config/v3"

	helmv1 "github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/v1"
)

func main() {
	helmPlugin := helmv1.New()
	c, err := cli.New(
		cli.WithPlugins(helmPlugin),
		cli.WithDefaultProjectVersion(configv3.Version),
		cli.WithDefaultPlugins(configv3.Version, helmPlugin),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
