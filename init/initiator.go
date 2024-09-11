package init

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const Name = "init"

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
	googleProtoFolder := "proto/google/"
	if err := os.MkdirAll(googleProtoFolder, 0744); err != nil {
		return fmt.Errorf("failed to create dir for dependency google proto files: %v", err)
	}

	repoURL := "https://github.com/googleapis/googleapis.git"
	cloneDir := "googleapis"
	protoDestDir := "./proto/google/"

	// Step 1: Clone the Google APIs repository
	if err := runCommand("git", "clone", repoURL); err != nil {
		return err
	}

	// Step 3: Copy necessary proto files
	protoFiles := []string{"annotations.proto", "http.proto", "httpbody.proto"}
	for _, protoFile := range protoFiles {
		srcPath := filepath.Join(cloneDir, "google/api", protoFile)
		destPath := filepath.Join(protoDestDir, protoFile)

		if err := copyFile(srcPath, destPath); err != nil {
			return err
		}
	}

	if err := os.RemoveAll(cloneDir); err != nil {
		return err
	}

	return nil
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = destFile.ReadFrom(sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
