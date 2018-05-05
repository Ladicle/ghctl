package repo

import (
	"io"

	"github.com/spf13/cobra"
)

// NewRepositoryCmd create new cobra command to handle repository.
func NewRepositoryCmd(out, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "repository [command]",
		Short:   "Manage repository data.",
		Aliases: []string{"repo"},
	}
	cmd.AddCommand(newGetCmd(out, errOut))
	cmd.AddCommand(newListCmd(out, errOut))
	return cmd
}
