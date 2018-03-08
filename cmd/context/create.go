package context

import (
	"io"

	"fmt"

	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/github"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

var defaultEndpoint = "https://api.github.com/graphql"

type createOption struct {
	ClientGenerator func(token string) *github.Client
	Token           string
	Endpoint        string
	Name            string
}

func newCreateCmd(out, errOut io.Writer) *cobra.Command {
	o := createOption{
		ClientGenerator: github.NewClient,
	}
	cmd := &cobra.Command{
		Use:   "create [options] <access_token>",
		Short: "Create new context",
		Run: func(cmd *cobra.Command, args []string) {
			util.HandleCmdError(o.validate(args), errOut)
			util.HandleCmdError(o.execute(out), errOut)
		},
	}
	cmd.Flags().StringVar(&o.Endpoint, "endpoint", defaultEndpoint, "GitHub API endpoint.")
	cmd.Flags().StringVar(&o.Name, "name", "", "Context identification.")
	return cmd
}

func (o *createOption) validate(args []string) error {
	if got, want := len(args), 1; got != want {
		return fmt.Errorf(
			"invalid number of arguments.",
			"want=%v, but got=%v",
			want, got)
	}
	o.Token = args[0]
	return nil
}

func (o *createOption) execute(out io.Writer) error {
	cli := o.ClientGenerator(o.Token)
	l, err := cli.GetLogin()
	if err != nil {
		return err
	}

	if o.Name == "" {
		o.Name = l
	}
	if err := config.RegisterContext(config.Context{
		Name:        o.Name,
		AccessToken: o.Token,
		Endpoint:    o.Endpoint,
	}); err != nil {
		return err
	}

	if err := config.SaveConfig(); err != nil {
		return err
	}
	fmt.Fprintf(out, "Register %s context is successfully.\n", o.Name)
	return nil
}
