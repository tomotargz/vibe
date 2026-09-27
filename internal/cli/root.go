// Package cli defines the vibe command-line interface.
package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// NewRootCmd returns the root command of vibe.
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "vibe",
		Short: "ToDoとメモを1つのノートとして扱うCLI",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}
}

// Execute runs the root command and exits with status 1 on error.
func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
