package chart

import (
	"path"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

type HelmIgnore struct {
	machinery.TemplateMixin

	ProjectName string
}

func (t *HelmIgnore) SetTemplateDefaults() error {
	t.Path = path.Join("charts", t.ProjectName, ".helmignore")
	t.TemplateBody = helmIgnoreTemplate
	return nil
}

const helmIgnoreTemplate = `# Patterns to ignore when building packages.
# This supports shell glob matching, relative path matching, and
# negation (prefixed with !). Only one pattern per line.
.DS_Store
# Common VCS dirs
.git/
.gitignore
.bzr/
.bzrignore
.hg/
.hgignore
.svn/
# Common backup files
*.swp
*.bak
*.tmp
*.orig
*~
# Various IDEs
.project
.idea/
*.tmproj
.vscode/
`
