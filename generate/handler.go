package generate

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/hsblhsn/kick/parse"
)

//go:embed templates/handler.go.tmpl
var handlerTemplateFile string

type HandlerTemplateValue struct {
	HasInputParam  bool
	HasErrorOutput bool
	HasMethod      bool
	NamingPrefix   string
	InputModel     string
	OutputModel    string
	Methods        []string
	Path           string
}

func Handlers(sigs []*parse.Signature) (string, error) {
	var stringBuilder strings.Builder
	for _, sig := range sigs {
		code, err := Handler(sig)
		if err != nil {
			return "", err
		}
		if _, err := stringBuilder.WriteString(code); err != nil {
			return "", err
		}
		stringBuilder.WriteString("\n\n")
	}
	return stringBuilder.String(), nil
}

func Handler(sig *parse.Signature) (string, error) {
	var (
		namingPrefix = sig.Pkg + "_" + sig.Func
		inputModel   = sig.Inputs[1]
		outputModel  = sig.Outputs[0]
		hasMethod    = len(sig.Methods) > 0
	)
	if isLocalIdent(inputModel) {
		inputModel = sig.Pkg + "." + inputModel
	}
	if isLocalIdent(outputModel) {
		outputModel = sig.Pkg + "." + outputModel
	}
	return Template(handlerTemplate, HandlerTemplateValue{
		HasInputParam:  len(sig.Inputs) > 1,
		HasErrorOutput: len(sig.Outputs) > 1,
		HasMethod:      hasMethod,
		NamingPrefix:   namingPrefix,
		InputModel:     inputModel,
		OutputModel:    outputModel,
		Methods:        sig.Methods,
		Path:           sig.Path,
	})
}

var (
	handlerTemplate = template.Must(template.New("handler").Parse(handlerTemplateFile))
)

func Template(tpl *template.Template, value interface{}) (string, error) {
	var stringBuilder strings.Builder
	if err := tpl.Execute(&stringBuilder, value); err != nil {
		return "", err
	}
	return stringBuilder.String(), nil
}

func isLocalIdent(ident string) bool {
	return isExported(ident) && !strings.Contains(ident, ".")
}

// isExported returns true if the first char of the ident is uppercase.
func isExported(ident string) bool {
	return ident[0] >= 'A' && ident[0] <= 'Z'
}
