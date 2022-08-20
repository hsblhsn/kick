package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type Signature struct {
	Pkg     string
	Func    string
	Inputs  []string
	Outputs []string
	Methods []string
	Path    string
}

// File loads the Go file and processes it.
func File(filename string) ([]*Signature, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	text := string(data)
	signatures, err := extractSignatures(filename, text)
	if err != nil {
		return nil, err
	}
	return signatures, nil
}

func extractSignatures(filename, content string) ([]*Signature, error) {
	fset := token.NewFileSet()
	var node *ast.File
	node, err := parser.ParseFile(fset, filename, content, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	signatures := make([]*Signature, 0)
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		lastLineOfDoc := funcDecl.Doc.List[len(funcDecl.Doc.List)-1].Text
		if !IsDirective(lastLineOfDoc) {
			continue
		}

		directive, err := Directive(lastLineOfDoc)
		if err != nil {
			return nil, err
		}

		sig, err := getFunctionSignature(funcDecl)
		if err != nil {
			return nil, err
		}
		sig.Pkg = node.Name.Name
		sig.Methods = directive.Methods
		sig.Path = directive.Path
		signatures = append(signatures, sig)
	}
	return signatures, nil
}

func getFunctionSignature(funcDecl *ast.FuncDecl) (*Signature, error) {
	funcName := getFuncName(funcDecl)
	inputs := getFieldSignatures(funcDecl.Type.Params.List)
	outputs := getFieldSignatures(funcDecl.Type.Results.List)
	return &Signature{
		Func:    funcName,
		Inputs:  inputs,
		Outputs: outputs,
	}, nil
}

func getFieldSignatures(fields []*ast.Field) []string {
	signatures := make([]string, len(fields))
	for i, field := range fields {
		sign := getParamSignature(field.Type)
		signatures[i] = sign
	}
	return signatures
}

func getFuncName(funcDecl *ast.FuncDecl) string {
	return funcDecl.Name.Name
}

func getParamSignature(expr ast.Expr) string {
	switch exp := expr.(type) {
	case *ast.Ident:
		return exp.Name
	case *ast.SelectorExpr:
		return getParamSignature(exp.X) + "." + getParamSignature(exp.Sel)
	default:
		return ""
	}
}
