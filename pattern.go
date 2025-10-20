package vanda

import (
	"fmt"
	"strings"
)

func ParsePattern(pattern string) (*CommandPattern, error) {
	parts := strings.Fields(pattern)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty pattern")
	}

	// Process subcommands
	if !strings.HasPrefix(parts[0], "<") || !strings.HasSuffix(parts[0], ">") {
		//TODO: when no subcommands are defined, just follow a single pattern
		return nil, fmt.Errorf("pattern must start with <command>")
	}

	cmdName := strings.Trim(parts[0], "<>")

	var args []ArgPattern

	var flags []string

	if cmdName != "" {
		flags = parts[1:]
	} else {
		flags = parts
	}

	for _, part := range flags {
		switch {
		// required argument
		case strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]"):
			content := strings.Trim(part, "[]")
			nameType := strings.Split(content, ":")
			if len(nameType) != 2 {
				return nil, fmt.Errorf("invalid required argument: %s", part)
			}
			args = append(args, ArgPattern{
				Name:     nameType[0],
				Type:     ArgType(nameType[1]),
				Required: true,
			})

		// optional arg
		case strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}"):
			content := strings.Trim(part, "{}")

			var nameType []string
			// flag with value: {=flag:type}
			if cntn, ok := strings.CutPrefix(content, "="); ok {
				nameType = strings.Split(cntn, ":")
				if len(nameType) != 2 {
					return nil, fmt.Errorf("invalid flag with value: %s", part)
				}
			} else {
				// boolean flag: {flag:bool}
				nameType = strings.Split(cntn, ":")
				if len(nameType) != 2 {
					return nil, fmt.Errorf("invalid flag: %s", part)
				}
			}

			args = append(args, ArgPattern{
				Name: nameType[0],
				Flag: "-" + nameType[0],
				Type: ArgType(nameType[1]),
			})

		}
	}

	return &CommandPattern{Name: cmdName, Args: args}, nil
}
