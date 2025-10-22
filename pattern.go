package vanda

import (
	"fmt"
	"strings"
)

func ParsePattern(pattern string) (*CommandPattern, error) {
	parts := strings.Fields(pattern)
	if len(parts) == 0 {
		return nil, &ArgumentParsingError{
			"Empty Pattern",
		}
	}

	// Process subcommands
	if !strings.HasPrefix(parts[0], "<") || !strings.HasSuffix(parts[0], ">") {
		//TODO: when no subcommands are defined, just follow a single pattern
		return nil, &ArgumentParsingError{
			"Pattern must start with <command>",
		}
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
				return nil, &ArgumentParsingError{
					fmt.Sprintf("Invalid required argument: %s", part),
				}
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
					return nil, &ArgumentParsingError{
						fmt.Sprintf("Invalid flag: %s", part),
					}
				}
				if cntn1, ok1 := strings.CutPrefix(nameType[0], "="); ok1 {
					println("with -- ", fmt.Sprintf("--%s", cntn1))
					args = append(args, ArgPattern{
						Name: cntn1,
						Flag: fmt.Sprintf("--%s", cntn1),
						Type: ArgType(nameType[1]),
					})
				} else {
					println("without --", nameType[0])
					args = append(args, ArgPattern{
						Name: nameType[0],
						Flag: nameType[0],
						Type: ArgType(nameType[1]),
					})
				}
			}
		}

	}
	return &CommandPattern{Name: cmdName, Args: args}, nil
}
