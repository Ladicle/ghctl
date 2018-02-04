package cmd

import (
	"io"
	"os"

	cctx "github.com/Ladicle/ghctl/cmd/context"
	"github.com/Ladicle/ghctl/pkg/config"
	"github.com/Ladicle/ghctl/pkg/util"
	"github.com/spf13/cobra"
)

var (
	cfgFile string

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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "ghconfig", "", "config file (default: $HOME/.ghctl/config)")
	rootCmd.AddCommand(NewVersionCmd(out, errOut))
	rootCmd.AddCommand(cctx.NewContextCmd(out, errOut))
	return rootCmd
}

// Execute runs ghctl command.
func Execute() {
	errOut := os.Stderr
	util.HandleCmdError(newRootCmd(os.Stdout, errOut).Execute(), errOut)
}

func initConfig() {
	errOut := os.Stderr
	if cfgFile != "" {
		util.HandleCmdError(config.SetConfigFile(cfgFile), errOut)
	} else {
		util.HandleCmdError(config.SetDefaultConfigFile(), errOut)
	}
	util.HandleCmdError(config.LoadConfig(), errOut)
}
