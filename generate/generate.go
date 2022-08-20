package generate

import (
	"github.com/hsblhsn/kick/parse"
)

func Generate(sigs []*parse.Signature) (string, error) {
	handlerCode, err := Handlers(sigs)
	if err != nil {
		return "", err
	}

	return handlerCode, nil
}
