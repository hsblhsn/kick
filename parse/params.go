package parse

import (
	"fmt"
	"strings"
)

var (
	ErrInvalidDirective = fmt.Errorf("invalid directive")
	ErrMissingDirective = fmt.Errorf("missing directive")
)

const (
	DirectiveMethod = "methods"
	DirectivePath   = "path"
)

type DirectiveParams map[string]string

func NewDirectiveParams(line string) (DirectiveParams, error) {
	params := make(DirectiveParams)
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return params, nil
	}
	for _, part := range parts {
		kv := strings.Split(part, "=")
		if len(kv) != 2 {
			return params, fmt.Errorf("invalid directive: %q: %w", line, ErrInvalidDirective)
		}
		params[kv[0]] = kv[1]
	}
	return params, nil
}

func (d DirectiveParams) Get(key string) (string, error) {
	if v, ok := d[key]; ok {
		return v, nil
	}
	return "", fmt.Errorf("missing directive: %q: %w", key, ErrMissingDirective)
}

func (d DirectiveParams) GetMethods() ([]string, error) {
	method, err := d.Get(DirectiveMethod)
	if err != nil {
		return nil, err
	}
	methods := strings.Split(method, ",")
	for i, m := range methods {
		methods[i] = strings.ToUpper(strings.TrimSpace(m))
	}
	return methods, nil
}

func (d DirectiveParams) GetPath() (string, error) {
	path, err := d.Get(DirectivePath)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(path, "/") {
		return path, nil
	}
	return "/" + path, nil
}
