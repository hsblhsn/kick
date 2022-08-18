package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kick",
		Short: "kick cli",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
