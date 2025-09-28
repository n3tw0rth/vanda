package vanda

import (
	"fmt"
	"strconv"
	"strings"
)

// ArgType represents the type of an argument
type ArgType string

const (
	String ArgType = "string"
	Int    ArgType = "int"
	Bool   ArgType = "bool"
)

// ArgPattern describes a single argument or flag
type ArgPattern struct {
	Name     string
	Flag     string
	Type     ArgType
	Required bool
}

// CommandPattern represents a command with its argument patterns
type CommandPattern struct {
	Name string
	Args []ArgPattern
}

// Parser holds multiple command patterns
type Parser struct {
	Commands []*CommandPattern
}

// NewParser creates a parser from a list of DSL patterns
func NewParser(patterns []string) (*Parser, error) {
	cmds := []*CommandPattern{}
	for _, pat := range patterns {
		cmd, err := ParsePattern(pat)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	return &Parser{Commands: cmds}, nil
}

// MatchAndParse picks the correct command pattern based on argv[0] and parses the rest
func (p *Parser) MatchAndParse(argv []string) (string, map[string]interface{}, error) {
	if len(argv) == 0 {
		return "", nil, fmt.Errorf("no command provided")
	}

	cmdArg := argv[0]
	var selected *CommandPattern
	for _, cmd := range p.Commands {
		if cmd.Name == cmdArg {
			selected = cmd
			break
		}
	}
	if selected == nil {
		return "", nil, fmt.Errorf("unknown command: %s", cmdArg)
	}

	// parse the remaining args
	result := make(map[string]interface{})
	rest := argv[1:]

	for _, pat := range selected.Args {
		if pat.Required {
			if len(rest) == 0 {
				return "", nil, fmt.Errorf("missing required argument: %s", pat.Name)
			}
			val, err := castValue(rest[0], pat.Type)
			if err != nil {
				return "", nil, err
			}
			result[pat.Name] = val
			rest = rest[1:]
		} else {
			// optional flags
			if pat.Type == Bool {
				for _, a := range rest {
					if a == "-"+pat.Name {
						result[pat.Name] = true
					}
				}
			} else {
				for i := 0; i < len(rest); i++ {
					if strings.HasPrefix(rest[i], "-"+pat.Name) {
						if rest[i] == "-"+pat.Name && i+1 < len(rest) {
							val, err := castValue(rest[i+1], pat.Type)
							if err != nil {
								return "", nil, err
							}
							result[pat.Name] = val
						} else {
							val, err := castValue(strings.TrimPrefix(rest[i], "-"+pat.Name), pat.Type)
							if err != nil {
								return "", nil, err
							}
							result[pat.Name] = val
						}
					}
				}
			}
		}
	}

	return selected.Name, result, nil
}

func castValue(s string, t ArgType) (interface{}, error) {
	switch t {
	case String:
		return s, nil
	case Int:
		return strconv.Atoi(s)
	case Bool:
		return strconv.ParseBool(s)
	default:
		return nil, fmt.Errorf("unsupported type: %s", t)
	}
}
