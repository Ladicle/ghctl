package repository

import (
	"fmt"
	"io"

	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/github"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

type listOption struct {
	ClientGenerator func(token string) *github.Client
	Organization    string
	Cache           bool
	Match           string
	Output          string
}

func newListCmd(out, errOut io.Writer) *cobra.Command {
	o := listOption{
		ClientGenerator: github.NewClient,
	}
	cmd := &cobra.Command{
		Use:   "list [options] <organization>",
		Short: "lists repositories",
		Run: func(cmd *cobra.Command, args []string) {
			util.HandleCmdError(o.validate(args), errOut)
			util.HandleCmdError(o.execute(out), errOut)
		},
	}

	cmd.Flags().BoolVar(&o.Cache, "cache", false, "Use cache.")
	cmd.Flags().StringVar(&o.Match, "match", "", "Show matched repository data.")
	cmd.Flags().StringVarP(&o.Output, "output", "o", "", "Output format.")
	return cmd
}

func (o *listOption) validate(args []string) error {
	if len(args) == 1 {
		o.Organization = args[0]
	} else {
		return fmt.Errorf("organization is required")
	}

	if o.Output != "yaml" && o.Output != "json" {
		return fmt.Errorf("%s is unknown output format", o.Output)
	}
	return nil
}

func (o *listOption) execute(out io.Writer) error {
	var org *github.Organization
	if o.Cache {
		// TODO: get data from cache
		return fmt.Errorf("cache option have not implemented yet")
	} else {
		cli := o.ClientGenerator(config.GetCurrentContext().AccessToken)
		newOrg, err := cli.GetOrganization(o.Organization)
		if err != nil {
			return err
		}
		org = newOrg

		// TODO: update cache
	}

	// TODO: filter repositories by match

	if d, err := util.GetPrettyOutput(o.Output, *org); err != nil {
		return err
	} else {
		fmt.Fprintf(out, "%s", string(d))
	}
	return nil
}
