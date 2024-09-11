package adder

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jarium/go-proto-cli/content"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

const Name = "add"

var (
	ErrNoProtoName = errors.New("no proto file name provided")
)

type Args struct {
	Name string
	HTTP bool
}

type Adder struct {
	ar *Args
}

func NewAdder() *Adder {
	return &Adder{ar: &Args{}}
}

func (a *Adder) GetName() string {
	return Name
}

func (a *Adder) SetArgs(set *flag.FlagSet) {
	set.StringVar(&a.ar.Name, "name", "", "Name of proto file")
	set.BoolVar(&a.ar.HTTP, "http", false, "Generate HTTP server")
}

func (a *Adder) Execute() error {
	if a.ar.Name == "" {
		return ErrNoProtoName
	}

	b := content.NewBuilder()

	b.Addln(`syntax = "proto3";`)
	b.Br(1)
	b.Addfln("package %s;", a.ar.Name)
	b.Br(1)
	b.Addfln(`option go_package = "./%s_pb";`, a.ar.Name)

	if a.ar.HTTP {
		b.Br(1)
		b.Addln(`import "google/annotations.proto";`)
	}

	b.Br(1)

	ucName := cases.Title(language.English).String(a.ar.Name)
	b.Addfln("// The %s service definition.", ucName)
	b.Addfln("service %s {", ucName)
	b.Addln("  rpc Example(ExampleRequest) returns (ExampleReply) {")

	if a.ar.HTTP {
		b.Addln("    option (google.api.http) = {")
		b.Addln(`      get: "/example"`)
		b.Addln(`      body: "*"`)
		b.Addln("    };")
	}

	b.Addln("  }")
	b.Addln("}")
	b.Br(1)
	b.Addln("message ExampleRequest {")
	b.Addln("  string name = 1;")
	b.Addln("}")
	b.Br(1)
	b.Addln("message ExampleReply {")
	b.Addln("  string name = 1;")
	b.Addln("}")

	folder := fmt.Sprintf("proto/gen/%s_pb/", a.ar.Name)
	filePath := fmt.Sprintf("%s%s.proto", folder, a.ar.Name)

	if err := os.MkdirAll(folder, 0744); err != nil {
		return fmt.Errorf("failed to create dir for .proto file: %v", err)
	}

	if err := os.WriteFile(filePath, []byte(b.Get()), 0644); err != nil {
		return fmt.Errorf("failed to write .proto file: %v", err)
	}

	return nil
}
