package context

import (
	"fmt"
	"io"

	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

type currentOption struct {
	SwitchCtxName string
	SimpleOutput  bool
}

func newCurrentCmd(out, errOut io.Writer) *cobra.Command {
	o := currentOption{}
	cmd := &cobra.Command{
		Use:   "current [options]",
		Short: "Displays the current context and switch it.",
		Run: func(cmd *cobra.Command, args []string) {
			util.HandleCmdError(o.execute(out), errOut)
		},
	}
	cmd.Flags().StringVar(&o.SwitchCtxName, "switch", "", "Switch current context.")
	cmd.Flags().BoolVar(&o.SimpleOutput, "simple", false, "Displays only current context name.")
	return cmd
}

func (o *currentOption) execute(out io.Writer) error {
	if o.SwitchCtxName != "" {
		if err := config.SetCurrentContext(o.SwitchCtxName); err != nil {
			return err
		}
		fmt.Fprintf(out, "The Current context switches to %s.\n", o.SwitchCtxName)
		return nil
	}
	name := config.GetCurrentContext()
	if o.SimpleOutput {
		fmt.Fprintf(out, "%s", name)
		return nil
	}
	fmt.Fprintf(out, "Current context name is %s.\n", name)
	return nil
}
