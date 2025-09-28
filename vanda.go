package vanda

import (
	"fmt"
	"strconv"
	"strings"
)

type ArgType string

const (
	String ArgType = "string"
	Int    ArgType = "int"
	Bool   ArgType = "bool"
)

type ArgPattern struct {
	Name     string
	Flag     string
	Type     ArgType
	Required bool
}

type Parser struct {
	pattern []ArgPattern
}

// New creates a parser from DSL pattern
func New(pattern string) (*Parser, error) {
	args, err := parsePattern(pattern)
	if err != nil {
		return nil, err
	}
	return &Parser{pattern: args}, nil
}

// Parse command-line args against the pattern
func (p *Parser) Parse(argv []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	argi := 0

	for _, pat := range p.pattern {
		if pat.Required {
			if argi >= len(argv) {
				return nil, fmt.Errorf("missing required argument: %s", pat.Name)
			}
			val, err := castValue(argv[argi], pat.Type)
			if err != nil {
				return nil, err
			}
			result[pat.Name] = val
			argi++
		} else {
			if pat.Type == Bool {
				for _, a := range argv {
					if a == "-"+pat.Name {
						result[pat.Name] = true
					}
				}
			} else {
				for i := 0; i < len(argv); i++ {
					if strings.HasPrefix(argv[i], "-"+pat.Name) {
						if argv[i] == "-"+pat.Name && i+1 < len(argv) {
							val, err := castValue(argv[i+1], pat.Type)
							if err != nil {
								return nil, err
							}
							result[pat.Name] = val
						} else {
							val, err := castValue(strings.TrimPrefix(argv[i], "-"+pat.Name), pat.Type)
							if err != nil {
								return nil, err
							}
							result[pat.Name] = val
						}
					}
				}
			}
		}
	}

	return result, nil
}

func castValue(s string, t ArgType) (interface{}, error) {
	switch t {
	case String:
		return s, nil
	case Int:
		return strconv.Atoi(s)
	case Bool:
		// presence already handled, but if value-style
		return strconv.ParseBool(s)
	default:
		return nil, fmt.Errorf("unsupported type: %s", t)
	}
}
