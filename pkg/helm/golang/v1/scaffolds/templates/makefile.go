package templates

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

type Makefile struct {
	machinery.TemplateMixin

	Image                    string
	EnvtestK8sVersion        string
	ControllerToolsVersion   string
	ControllerRuntimeVersion string
	HelmVersion              string
}

func (t *Makefile) SetTemplateDefaults() error {
	t.Path = "Makefile"
	t.Image = "controller:latest"
	t.TemplateBody = makefileTemplate
	return nil
}

// Based on https://github.com/kubernetes-sigs/kubebuilder/blob/master/pkg/plugins/golang/v3/scaffolds/internal/templates/makefile.go.
const makefileTemplate = `# Image URL to use all building/pushing image targets.
IMG ?= {{ .Image }}
# Kubernetes version to use by envtest in tests.
ENVTEST_K8S_VERSION ?= {{ .EnvtestK8sVersion }}
# controller-tools version, used for code and manifest generation.
CONTROLLER_TOOLS_VERSION ?= {{ .ControllerToolsVersion }}
# controller-runtime version.
CONTROLLER_RUNTIME_VERSION ?= {{ .ControllerRuntimeVersion }}
# Helm version to use when deploying.
HELM_VERSION ?= {{ .HelmVersion }}

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

generate: controller-gen ## Generate DeepCopy code and Helm chart.
	$(CONTROLLER_GEN) object paths="./..."

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

test: generate fmt vet envtest ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path)" go test ./... -coverprofile cover.out

##@ Build

build: generate fmt vet ## Build manager binary.
	go build -o bin/manager main.go

run: generate fmt vet ## Run a controller from your host.
	go run ./main.go

docker-build: test ## Build docker image with the manager.
	docker build -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}

##@ Deployment

install: generate helm ## Deploy Helm chart to the K8s cluster specified in ~/.kube/config.
	$(HELM) install --upgrade

uninstall: generate helm ## Uninstall Helm chart from the K8s cluster specified in ~/.kube/config.
	$(HELM) uninstall


CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION))

HELM = $(shell pwd)/bin/helm
helm: ## Download helm locally if necessary.
	$(call go-get-tool,$(HELM),helm.sh/helm/v3@$(HELM_VERSION))

ENVTEST = $(shell pwd)/bin/setup-envtest
envtest: ## Download envtest-setup locally if necessary.
	$(call go-get-tool,$(ENVTEST),sigs.k8s.io/controller-runtime/tools/setup-envtest@$(CONTROLLER_RUNTIME_VERSION))

# go-get-tool will 'go install' any package $2 with bin filepath $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
}
endef
`
