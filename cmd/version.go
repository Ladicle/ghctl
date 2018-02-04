package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// NewVersionCmd return cobra.Command to print a version information.
func NewVersionCmd(out, errOut io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Display version information of ghctl",
		Run: func(cmd *cobra.Command, args []string) {
			if version == "" {
				fmt.Fprintln(errOut, "unknown version")
				return
			}
			fmt.Fprintf(out, "ghctl version %v\n", version)
		},
	}
}
