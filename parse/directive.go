package parse

import "strings"

type DirectiveValue struct {
	Methods []string
	Path    string
}

// IsDirective returns true if the line is a directive.
func IsDirective(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "// kick: ")
}

// Directive parses a directive line from a kick directive.
// An example directive line is:
// kick: method=GET path=/hello-world
func Directive(line string) (*DirectiveValue, error) {
	line = strings.TrimPrefix(strings.TrimSpace(line), "// kick: ")
	params, err := NewDirectiveParams(line)
	if err != nil {
		return nil, err
	}
	methods, _ := params.GetMethods()
	path, err := params.GetPath()
	if err != nil {
		return nil, err
	}
	return &DirectiveValue{
		Methods: methods,
		Path:    path,
	}, nil
}
