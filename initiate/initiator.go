package initiate

import (
	"flag"
	"fmt"
	"github.com/jarium/go-proto-cli/executor"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
)

const Name = "initiate"

type Initiator struct {
}

func NewInitiator() *Initiator {
	return &Initiator{}
}

func (i *Initiator) GetName() string {
	return Name
}

func (i *Initiator) SetArgs(*flag.FlagSet) {
}

func (i *Initiator) Execute() error {
	httpPlugin := "github.com/jarium/protoc-gen-http"

	if err := executor.Exec("go", "get", httpPlugin); err != nil {
		return err
	}
	if err := executor.Exec("go", "install", httpPlugin); err != nil {
		return err
	}

	if err := os.MkdirAll("proto/google/", 0744); err != nil {
		return fmt.Errorf("failed to create dir for dependency google proto files: %v", err)
	}

	modPath, err := findModulePath("github.com/jarium/go-proto-cli")
	if err != nil {
		return fmt.Errorf("error finding module path: %w", err)
	}

	srcFolder := filepath.Join(modPath, "initiate/google/")
	srcFiles, err := os.ReadDir(srcFolder)

	if err != nil {
		return fmt.Errorf("error when reading proto files under google folder: %w", err)
	}

	for _, f := range srcFiles {
		if err := copyFile(filepath.Join(srcFolder, f.Name()), "proto/google"); err != nil {
			return fmt.Errorf("error when copying file: %s, err: %w", f.Name(), err)
		}
	}

	return nil
}

func copyFile(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// findModulePath uses Go's module cache to locate the path to an external library.
func findModulePath(moduleName string) (string, error) {
	modInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("failed to read build info")
	}

	for _, dep := range modInfo.Deps {
		if dep.Path == moduleName {
			return dep.Replace.Path, nil
		}
	}

	return "", fmt.Errorf("module %s not found in dependencies", moduleName)
}
