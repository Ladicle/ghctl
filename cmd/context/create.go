package user

import (
	"context"
	"io"

	"fmt"

	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/github"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

var defaultEndpoint = "https://api.github.com/graphql"

type createOption struct {
	Token    string
	Endpoint string
	Name     string
}

func newCreateCmd(out, errOut io.Writer) *cobra.Command {
	c := createOption{}
	createCmd := &cobra.Command{
		Use:   "create [options...] <access_token>",
		Short: "Create new context",
		Run: func(cmd *cobra.Command, args []string) {
			util.HandleCmdError(c.validate(args), errOut)
			util.HandleCmdError(c.execute(out), errOut)
		},
	}
	createCmd.Flags().StringVar(&c.Endpoint, "endpoint", defaultEndpoint, "GitHub API endpoint.")
	createCmd.Flags().StringVar(&c.Name, "name", "", "Context identification.")
	return createCmd
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
	q := github.LoginQuery{}
	cli := github.NewClient(o.Token)
	if err := cli.Query(context.Background(), &q, nil); err != nil {
		return err
	}

	if o.Name == "" {
		o.Name = string(q.Viewer.Login)
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
