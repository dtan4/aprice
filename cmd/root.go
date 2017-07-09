// Copyright © 2017 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	debug bool
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

	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
