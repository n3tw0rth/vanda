package main

import (
	"fmt"
	"github.com/n3tw0rth/vanda"
)

func main() {
	patterns := []string{
		"<start> [integration:string] [key:string] {=from:string} {=to:string}",
		"<nmap> [ip:string] {sV:bool} {=oN:string} {=minrate:int}",
	}
	argv := []string{"start", "jira", "knox", "-from", "1000", "-to", "1030"}
	// argv := []string{"nmap", "10.10.10.10", "-sV", "-oN", "scan.txt", "-minrate", "1000"}

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
