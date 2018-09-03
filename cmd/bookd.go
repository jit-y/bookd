package cmd

import "github.com/spf13/cobra"

func NewRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bookd",
		Short: "",
		Long:  "",
	}

	return cmd
}
