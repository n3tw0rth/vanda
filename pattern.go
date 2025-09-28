package vanda

import (
	"fmt"
	"strings"
)

func parsePattern(pattern string) ([]ArgPattern, error) {
	parts := strings.Fields(pattern)
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty pattern")
	}
	var args []ArgPattern

	for _, part := range parts[1:] { // skip command name
		if strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]") {
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
		} else if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			content := strings.Trim(part, "{}")
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
	return args, nil
}
