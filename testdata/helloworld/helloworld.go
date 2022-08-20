package helloworld

import (
	"context"
	"fmt"
)

// HelloWorldParams is the input to the HelloWorld function.
type HelloWorldParams struct {
	Name string
}

// HelloWorldResponse is the response type for the HelloWorld function.
type HelloWorldResponse struct {
	Message string
}

// HelloWorld is a simple function that returns a string.
// kick: methods=GET path=/hello-world
func HelloWorld(ctx context.Context, params HelloWorldParams) (HelloWorldResponse, error) {
	if params.Name == "" {
		return HelloWorldResponse{Message: "Hello, World!"}, nil
	}
	return HelloWorldResponse{
		Message: fmt.Sprintf("Hello, %s!", params.Name),
	}, nil
}
