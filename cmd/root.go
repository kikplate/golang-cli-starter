package cmd

import (
	"fmt"
	"os"

	"github.com/kikplate/golang-cli-starter/internal/config"
	"github.com/kikplate/golang-cli-starter/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cliforge",
		Short: "Forge your CLI tools with confidence",
		Long: `Cliforge is a production-ready CLI boilerplate built with Cobra.
It demonstrates config management, structured logging, subcommands, and testing.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logger.Init(verbose)
			return config.Load(cfgFile)
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: $HOME/.cliforge.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newGreetCmd())

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
