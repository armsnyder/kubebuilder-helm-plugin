package test

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	"github.com/pmezard/go-difflib/difflib"
)

// runCommand runs the compiled kubebuilderhelm command.
var runCommand func(ops ...commandOption) error

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "plugin suite")
}

var _ = BeforeSuite(func() {
	// Show full diff, useful when comparing scaffolded file contents.
	format.TruncatedDiff = false

	// Calculate an absolute path to the kubebuilderhelm binary, since we will be running it from
	// inside test directories.
	_, callerFilename, _, _ := runtime.Caller(0)
	absBinPath := filepath.Join(filepath.Dir(callerFilename), "..", "bin", "kubebuilderhelm")

	// Compile kubebuilderhelm.
	goBuildCmd := exec.Command("go", "build", "-o", absBinPath, filepath.Join("..", "cmd", "kubebuilderhelm"))
	goBuildCmd.Stdout = GinkgoWriter
	goBuildCmd.Stderr = GinkgoWriter
	Expect(goBuildCmd.Run()).To(Succeed())

	runCommand = func(ops ...commandOption) error {
		command := exec.Command(absBinPath)
		command.Stdout = GinkgoWriter
		command.Stderr = GinkgoWriter
		for _, op := range ops {
			op(command)
		}
		By("running " + command.String())
		return command.Run()
	}
})

type commandOption func(cmd *exec.Cmd)

func withArgs(args ...string) commandOption {
	return func(cmd *exec.Cmd) {
		cmd.Args = append([]string{cmd.Path}, args...)
	}
}

func inDir(dir string) commandOption {
	return func(cmd *exec.Cmd) {
		cmd.Dir = dir
	}
}

// makeTempDir creates a temporary directory in the OS's default temporary directory. The resulting
// directory name will patch the provided name exactly. The caller is expected to call the returned
// cleanup function to remove the directory after use.
func makeTempDir(name string) (dir string, rm func()) {
	tempDir, err := os.MkdirTemp("", "kubebuilder-helm-plugin-"+name)
	Expect(err).NotTo(HaveOccurred())

	dir = filepath.Join(tempDir, name)
	Expect(os.MkdirAll(dir, 0755)).To(Succeed())

	rm = func() {
		Expect(os.RemoveAll(tempDir)).To(Succeed())
	}

	return dir, rm
}

// shouldScaffoldFiles may be passed to Gomega's It() in order to assert that the files and file
// contents match exactly between the provided testdata subdirectory name and an actual directory.
func shouldScaffoldFiles(expectedTestDataSubDirName string, actualDir *string) (string, func()) {
	return "should scaffold files", func() {
		expectScaffoldedFilesMatch(expectedTestDataSubDirName, *actualDir)
	}
}

// expectScaffoldedFilesMatch asserts that the files and file contents match exactly between the
// provided testdata subdirectory name and an actual directory.
func expectScaffoldedFilesMatch(expectedTestDataSubDirName, actualDir string) {
	type FilePath struct {
		Path  string
		IsDir bool
	}

	// First check that the every file was scaffolded, regardless of content.
	explorePaths := func(root string) []FilePath {
		var result []FilePath
		Expect(fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
			Expect(err).NotTo(HaveOccurred())
			result = append(result, FilePath{Path: path, IsDir: d.IsDir()})
			return nil
		})).To(Succeed())
		return result
	}
	expectedDir := filepath.Join("testdata", expectedTestDataSubDirName)
	expectedPaths := explorePaths(expectedDir)
	actualPaths := explorePaths(actualDir)
	Expect(actualPaths).To(ConsistOf(expectedPaths), "scaffolded file paths")

	// Next check the file contents.
	for _, filePath := range expectedPaths {
		if filePath.IsDir {
			continue
		}
		expectedBytes, err := os.ReadFile(filepath.Join(expectedDir, filePath.Path))
		Expect(err).NotTo(HaveOccurred())
		actualBytes, err := os.ReadFile(filepath.Join(actualDir, filePath.Path))
		Expect(err).NotTo(HaveOccurred())
		Expect(string(actualBytes)).To(equalMultiline(string(expectedBytes)), "contents of scaffolded file %s", filePath.Path)
	}
}

// equalMultiline behaves like Gomega's Equal() matcher but prints diffs instead of truncated values
// on failure.
func equalMultiline(expected interface{}) types.GomegaMatcher {
	return &equalMultilineMatcher{EqualMatcher: matchers.EqualMatcher{Expected: expected}}
}

type equalMultilineMatcher struct {
	matchers.EqualMatcher
}

func (m *equalMultilineMatcher) FailureMessage(actual interface{}) (message string) {
	actualString, actualOK := actual.(string)
	expectedString, expectedOK := m.Expected.(string)

	// Use the standard formatter if we are not matching strings.
	if !(actualOK && expectedOK) {
		return format.Message(actual, "to equal", m.Expected)
	}

	// Otherwise, return a diff.
	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(expectedString),
		B:        difflib.SplitLines(actualString),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})

	return "Diff:\n" + diff
}
