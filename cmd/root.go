package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const (
	defaultApriceDirBase = ".aprice"
)

var (
	apriceDir string
	debug     bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "aprice",
	Short: "AWS Price List CLI",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if debug {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&apriceDir, "aprice-dir", "", "aprice data directory (default: ~/.aprice)")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if apriceDir == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		apriceDir = filepath.Join(home, defaultApriceDirBase)
	} else {
		expanded, err := homedir.Expand(apriceDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		abs, err := filepath.Abs(expanded)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		apriceDir = abs
	}
}
