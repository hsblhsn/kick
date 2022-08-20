package parse

import (
	"errors"
	"reflect"
	"testing"
)

func TestDirective(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name        string
		args        args
		want        *DirectiveValue
		expectedErr error
	}{
		{
			name: "valid directive",
			args: args{
				line: "// kick: method=GET path=/hello-world",
			},
			want: &DirectiveValue{
				Methods: []string{"GET"},
				Path:    "/hello-world",
			},
			expectedErr: nil,
		},
		{
			name: "valid directive with multiple methods",
			args: args{
				line: "// kick: method=GET,POST,PUT path=/hello-world",
			},
			want: &DirectiveValue{
				Methods: []string{"GET", "POST", "PUT"},
				Path:    "/hello-world",
			},
			expectedErr: nil,
		},
		{
			name: "missing method",
			args: args{
				line: "// kick: path=/hello-world",
			},
			expectedErr: nil,
		},
		{
			name: "missing path",
			args: args{
				line: "// kick: method=GET",
			},
			expectedErr: ErrMissingDirective,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Directive(tt.args.line)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Directive() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Directive() = %v, want %v", got, tt.want)
			}
		})
	}
}
