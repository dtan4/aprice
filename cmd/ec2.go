package cmd

import (
	"github.com/spf13/cobra"
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:    "ec2",
	Short:  "Manage EC2 price list",
	Hidden: true,
}

func init() {
	RootCmd.AddCommand(ec2Cmd)
}
