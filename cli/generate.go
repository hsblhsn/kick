package cli

import (
	"github.com/spf13/cobra"
)

func NewGenerateCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate the application code",
		Long:  "Generate the application code",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
