package main

import (
	"fmt"
	"github.com/n3tw0rth/vanda"
)

func main() {
	patterns := []string{
		"<subcommand> [requiredone:string] [requiredtwo:string] {=optionalone:string} {=optionaltwo:string}",
	}
	argv := []string{"subcommand", "requiredparam1", "requiredparam2", "-optionalone", "10", "-optionaltwo", "20"}

	parser, err := vanda.NewParser(patterns)
	if err != nil {
		panic(err)
	}

	cmdName, args, err := parser.MatchAndParse(argv)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Command: %s\n", cmdName)
	fmt.Printf("Parsed args: %+v\n", args)
}
