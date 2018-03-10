package context

import (
	"fmt"
	"io"

	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

type getOption struct {
	Output      string
	ContextName string
}

func newGetCmd(out, errOut io.Writer) *cobra.Command {
	o := getOption{}
	cmd := &cobra.Command{
		Use:   "get [options] [context_name]",
		Short: "Outputs context data",
		Run: func(cmd *cobra.Command, args []string) {
			util.HandleCmdError(o.validate(args), errOut)
			util.HandleCmdError(o.execute(out), errOut)
		},
	}
	cmd.Flags().StringVar(&o.Output, "output", "yaml", "Output format.")
	return cmd
}

func (o *getOption) validate(args []string) error {
	if o.Output != "yaml" && o.Output != "json" {
		return fmt.Errorf("%s is unknown output format", o.Output)
	}
	switch len(args) {
	case 0:
		// Do not anything.
	case 1:
		o.ContextName = args[0]
	default:
		return fmt.Errorf("invalid number of arguments. want=0|1, got=%d", len(args))
	}
	return nil
}

func (o *getOption) execute(out io.Writer) error {
	var ctx interface{}
	if o.ContextName == "" {
		ctx = config.GetContexts()
	} else {
		ctx = config.GetContext(o.ContextName)
		if ctx.(*config.Context) == nil {
			return fmt.Errorf("%s is not exists", o.ContextName)
		}
	}
	if d, err := util.GetPrettyOutput(o.Output, ctx); err != nil {
		return err
	} else {
		fmt.Fprintf(out, "%s", string(d))
	}
	return nil
}
