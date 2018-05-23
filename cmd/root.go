package cmd

import (
	"io"
	"os"

	"github.com/Ladicle/ghctl/cmd/ctx"
	"github.com/Ladicle/ghctl/cmd/repo"
	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

var (
	cfgDir string

	// These variables are set at the build time.
	version string
	gitRepo string
)

func newRootCmd(out, errOut io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ghctl",
		Short: "Read GitHub notifications, issues and pull-requests.",
		Long:  "The ghctl CLI tool for GitHub notifications, issue and pull-requests.",
	}
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgDir, "ghconfig", "", "config directory (default: $HOME/.ghctl)")
	rootCmd.AddCommand(NewVersionCmd(out, errOut))
	rootCmd.AddCommand(ctx.NewContextCmd(out, errOut))
	rootCmd.AddCommand(repo.NewRepositoryCmd(out, errOut))
	return rootCmd
}

// Execute runs ghctl command.
func Execute() {
	errOut := os.Stderr
	util.HandleCmdError(newRootCmd(os.Stdout, errOut).Execute(), errOut)
}

func initConfig() {
	errOut := os.Stderr
	if cfgDir != "" {
		util.HandleCmdError(config.SetConfigDir(cfgDir), errOut)
	} else {
		util.HandleCmdError(config.SetDefaultConfigDir(), errOut)
	}
	util.HandleCmdError(config.LoadConfig(), errOut)
}
