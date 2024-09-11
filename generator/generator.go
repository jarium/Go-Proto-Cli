package generator

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jarium/go-proto-cli/executor"
	"strings"
)

const Name = "generate"

var (
	ErrNoFileName = errors.New("no proto file name provided")
)

type Args struct {
	Name string
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
	set.StringVar(&g.ar.Name, "name", "", "Name of proto file")
}

func (g *Generator) Execute() error {
	if g.ar.Name == "" {
		return ErrNoFileName
	}

	if !strings.HasSuffix(g.ar.Name, ".proto") {
		g.ar.Name += ".proto"
	}

	genFolder := "proto/gen/"
	folder := fmt.Sprintf("%s%s_pb/", genFolder, g.ar.Name[:len(g.ar.Name)-6])

	outArgs := []string{
		fmt.Sprintf("--go_out=%s", genFolder),
		fmt.Sprintf("--go-grpc_out=%s", genFolder),
		fmt.Sprintf("--http_out=%s", genFolder),
		"--proto_path=./proto",
		"--proto_path=./proto/google",
	}

	if err := executor.Exec("protoc", append(outArgs, folder+g.ar.Name)...); err != nil {
		return fmt.Errorf("failed to generate code from proto: %v", err)
	}

	return nil
}
