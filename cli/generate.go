package cli

import (
	"github.com/hsblhsn/kick/generate"
	"github.com/hsblhsn/kick/parse"
	"github.com/spf13/cobra"
)

func NewGenerateCMD() *cobra.Command {
	var (
		input, output string
	)
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate the application code",
		Long:  "Generate the application code",
		RunE: func(cmd *cobra.Command, args []string) error {
			sigs, err := parse.File(input)
			if err != nil {
				return err
			}
			code, err := generate.Generate(sigs)
			if err != nil {
				return err
			}
			return writeFile(output, code)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&input, "input", "i", "", "input file")
	flags.StringVarP(&output, "output", "o", "", "output file")
	return cmd
}

func writeFile(_ string, content string) error {
	print(content)
	return nil
}
