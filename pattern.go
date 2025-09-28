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

	// first part must be <command>
	if !strings.HasPrefix(parts[0], "<") || !strings.HasSuffix(parts[0], ">") {
		return nil, fmt.Errorf("pattern must start with <command>")
	}
	cmdName := strings.Trim(parts[0], "<>")

	var args []ArgPattern

	for _, part := range parts[1:] {
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

			// flag with value: {=flag:type}
			if strings.HasPrefix(content, "=") {
				content = strings.TrimPrefix(content, "=")
				nameType := strings.Split(content, ":")
				if len(nameType) != 2 {
					return nil, fmt.Errorf("invalid flag with value: %s", part)
				}
				args = append(args, ArgPattern{
					Name: nameType[0],
					Flag: "-" + nameType[0],
					Type: ArgType(nameType[1]),
				})
			} else {
				// boolean flag: {flag:bool}
				nameType := strings.Split(content, ":")
				if len(nameType) != 2 {
					return nil, fmt.Errorf("invalid flag: %s", part)
				}
				args = append(args, ArgPattern{
					Name: nameType[0],
					Flag: "-" + nameType[0],
					Type: ArgType(nameType[1]),
				})
			}
		}
	}

	return &CommandPattern{Name: cmdName, Args: args}, nil
}
