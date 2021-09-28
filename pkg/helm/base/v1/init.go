package v1

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/validation"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/armsnyder/kubebuilder-helm-plugin/pkg/helm/base/v1/scaffolds"
)

type initSubcommand struct {
	config config.Config
	name   string
}

func (s *initSubcommand) BindFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.name, "project-name", "", "name of this project")
}

func (s *initSubcommand) InjectConfig(c config.Config) error {
	s.config = c

	if s.name == "" {
		dir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %v", err)
		}
		s.name = strings.ToLower(filepath.Base(dir))
	}

	if errs := validation.IsDNS1123Label(s.name); len(errs) > 0 {
		return fmt.Errorf("project name %q is invalid: %v", s.name, errs)
	}

	return s.config.SetProjectName(s.name)
}

func (s *initSubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := &scaffolds.InitScaffolder{Config: s.config}
	scaffolder.InjectFS(fs)
	return scaffolder.Scaffold()
}
