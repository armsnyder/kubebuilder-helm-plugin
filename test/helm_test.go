package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm"
)

const helmPluginKey = helm.Domain + "/v1-alpha"

var _ = Describe(helmPluginKey, func() {
	// Override the suite-level withArgs option to include an argument for the plugin being tested.
	withArgs := func(args ...string) commandOption {
		fullArgs := append(args, "--plugins", helmPluginKey)
		return withArgs(fullArgs...)
	}

	Context("inside foo", func() {
		var (
			dir   string
			rmDir func()
		)

		BeforeEach(func() {
			dir, rmDir = makeTempDir("foo")
		})

		AfterEach(func() {
			rmDir()
		})

		When("run init command", func() {
			BeforeEach(func() {
				Expect(runCommand(inDir(dir), withArgs("init", "--project-name", "basic"))).To(Succeed())
			})

			It(shouldScaffoldFiles("basic", &dir))
		})
	})
})
