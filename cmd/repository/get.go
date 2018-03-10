package repository

import (
	"fmt"
	"io"
	"strings"

	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/github"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

type getOption struct {
	ClientGenerator func(token string) *github.Client
	Organization    string
	Repository      string
	Cache           bool
	Output          string
}

func newGetCmd(out, errOut io.Writer) *cobra.Command {
	o := getOption{
		ClientGenerator: github.NewClient,
	}
	cmd := &cobra.Command{
		Use:   "get [options] <org/repository>",
		Short: "Outputs repository data",
		Run: func(cmd *cobra.Command, args []string) {
			util.HandleCmdError(o.validate(args), errOut)
			util.HandleCmdError(o.execute(out), errOut)
		},
	}

	cmd.Flags().BoolVar(&o.Cache, "cache", false, "Use cache.")
	cmd.Flags().StringVarP(&o.Output, "output", "o", "yaml", "Output format.")
	return cmd
}

func (o *getOption) validate(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("target repository is required")
	}

	orgRepo := strings.Split(args[0], "/")
	if len(orgRepo) == 2 {
		o.Organization, o.Repository = orgRepo[0], orgRepo[1]
	} else {
		return fmt.Errorf("%s is not right 'org/repo' format", args[0])
	}

	if o.Output != "yaml" && o.Output != "json" {
		return fmt.Errorf("%s is unknown output format", o.Output)
	}
	return nil
}

func (o *getOption) execute(out io.Writer) error {
	var repo *github.Repository
	if o.Cache {
		// TODO: get data from cache
		return fmt.Errorf("cache option have not implemented yet")
	} else {
		cli := o.ClientGenerator(config.GetCurrentContext().AccessToken)
		newRepo, err := cli.GetRepository(o.Organization, o.Repository)
		if err != nil {
			return err
		}
		repo = newRepo

		// TODO: update cache
	}

	if d, err := util.GetPrettyOutput(o.Output, *repo); err != nil {
		return err
	} else {
		fmt.Fprintf(out, "%s", string(d))
	}
	return nil
}
