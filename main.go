package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jarium/go-proto-cli/adder"
	"github.com/jarium/go-proto-cli/generator"
	"os"
)

type ICommand interface {
	GetName() string
	SetArgs(set *flag.FlagSet)
	Execute() error
}

var Commands = map[string]ICommand{
	adder.Name:     adder.NewAdder(),
	generator.Name: generator.NewGenerator(),
}

var (
	ErrCommandRequired = errors.New("command required")
	ErrInvalidCommand  = errors.New("invalid command")
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println(ErrCommandRequired)
		os.Exit(1)
	}

	command := args[1]
	c, ok := Commands[command]

	if !ok {
		var commands string
		for n := range Commands {
			commands += n + ","
		}

		commands = commands[:len(commands)-1]
		fmt.Println(fmt.Errorf("%w, valid commands: %s", ErrInvalidCommand, commands))
		os.Exit(1)
	}

	set := flag.NewFlagSet(c.GetName(), flag.ExitOnError)
	c.SetArgs(set)

	if err := set.Parse(args[2:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("successfully executed command: %s\n", c.GetName())
}
