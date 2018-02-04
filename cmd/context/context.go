package user

import (
	"io"

	"github.com/spf13/cobra"
)

// NewContextCmd create new cobra command to handle contexts.
func NewContextCmd(out, errOut io.Writer) *cobra.Command {
	userCmd := &cobra.Command{
		Use:     "context [OPTION] [COMMAND]",
		Short:   "Manage context data for requesting GitHub API.",
		Aliases: []string{"ctx"},
	}
	userCmd.AddCommand(newCreateCmd(out, errOut))
	return userCmd
}
