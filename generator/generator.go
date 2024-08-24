package generator

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const Name = "generate"

var (
	ErrNoFileName = errors.New("no proto file name provided")
)

type Args struct {
	File string
}

type Generator struct {
	ar *Args
}

func NewGenerator() *Generator {
	return &Generator{ar: &Args{}}
}

func (g *Generator) GetName() string {
	return Name
}

func (g *Generator) SetArgs(set *flag.FlagSet) {
	set.StringVar(&g.ar.File, "file", "", "name of proto file")
}

func (g *Generator) Execute() error {
	if g.ar.File == "" {
		return ErrNoFileName
	}

	if !strings.HasSuffix(g.ar.File, ".proto") {
		g.ar.File += ".proto"
	}

	genFolder := "proto/gen/"
	folder := fmt.Sprintf("%s%s_pb/", genFolder, g.ar.File[:len(g.ar.File)-6])

	outArgs := []string{
		fmt.Sprintf("--go_out=%s", genFolder),
		fmt.Sprintf("--go-grpc_out=%s", genFolder),
		fmt.Sprintf("--go-http_out=%s", genFolder),
		"--proto_path=./proto",
		"--proto_path=./proto/google",
	}

	cmd := exec.Command("protoc", append(outArgs, folder+g.ar.File)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate code from proto: %v", err)
	}

	return nil
}
