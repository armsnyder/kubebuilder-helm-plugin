# Contributing to the Kubebuilder Helm Plugin

Hey, thanks for contributing to this project! :smile: :tada: I will make an effort to review your pull request as long as the GitHub actions are passing and you generally follow the guidelines here.

## Guidelines

The vision of this project is to be an alternative to the [Kustomize plugin](https://github.com/kubernetes-sigs/kubebuilder/tree/v3.1.0/pkg/plugins/common/kustomize/v1) that comes default with [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder). So, any improvements to the plugin that improve the experience of developing Helm charts for Kustomize projects will likely be accepted.

Before starting on a pull request, please make sure that there is an [issue](https://github.com/armsnyder/kubebuilder-helm-plugin/issues) for the proposed change, and comment that you plan on working on it. If one does not exist, please [create a new issue](https://github.com/armsnyder/kubebuilder-helm-plugin/issues/new/choose).

## Developing the plugin

Run `make` before committing in order to make sure that the code is tested and linted.

The tests in the `test` directory. They are end-to-end tests that invoke a Kubebuilder CLI that includes this plugin. They are written in BDD-style using the [Ginkgo](https://github.com/onsi/ginkgo) library. If your change causes a test to fail, it is most likely that the scaffolded output has changes. Please find the mentioned file in `test/testdata` and update it to match the actual scaffolded output in order to get the tests to pass.
