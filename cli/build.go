package cli

import (
	"github.com/spf13/cobra"
)

func NewBuildCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the application",
		Long:  "Build the application",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
