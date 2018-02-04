package context

import (
	"io"

	"github.com/spf13/cobra"
)

// NewContextCmd create new cobra command to handle contexts.
func NewContextCmd(out, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "context [OPTION] [COMMAND]",
		Short:   "Manage context data for requesting GitHub API.",
		Aliases: []string{"ctx"},
	}
	cmd.AddCommand(newCreateCmd(out, errOut))
	cmd.AddCommand(newGetCmd(out, errOut))
	return cmd
}
