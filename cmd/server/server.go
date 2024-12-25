package server

import (
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "start server",
		Long:    "start iniscan server",
		Example: "iniscan server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() error {
	return nil
}

func run() error {
	return nil
}
