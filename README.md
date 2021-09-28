# Kubebuilder Helm Plugin

[![](https://github.com/armsnyder/kubebuilder-helm-plugin/actions/workflows/main.yml/badge.svg)](https://github.com/armsnyder/kubebuilder-helm-plugin/actions/workflows/main.yml)

**Note:** This project is alpha and not intended for production use. It is likely missing features.

`helm.kubebuilder.armsnyder.com` is a config plugin for Kubebuilder that scaffolds a Helm chart for
an operator project. This is an alternative to the default `kustomize.common.kubebuilder.io` plugin.

## Usage

Kubebuilder version 3 only supports plugins compiled into the CLI. External plugin support is coming
in a future Kubebuilder release. To use the `helm.kubebuilder.armsnyder.com` plugin, you can use the
included `kubebuilderhelm` command or compile your own CLI.

### Using kubebuilderhelm

To install the latest version of `kubebuilderhelm` into your GOBIN:

```
go install github.com/armsnyder/kubebuilder-helm-plugin/cmd/kubebuilderhelm@latest
```

Then use `kubebuilderhelm` in place of `kubebuilder`.

```
kubebuilderhelm init
```

### Custom CLI Example

For full flexibility of which plugins are included in the CLI, you can build your own CLI and
include this plugin.

```go
package main

import (
	"log"

	helmv1 "github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/v1"
	"sigs.k8s.io/kubebuilder/v3/pkg/cli"
)

func main() {
	c, err := cli.New(cli.WithPlugins(helmv1.New()))
	if err != nil {
		log.Fatal(err)
	}
	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
```
