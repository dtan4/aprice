package cmd

import (
	"github.com/spf13/cobra"
)

// ec2InitCmd represents the ec2 init command
var ec2InitCmd = &cobra.Command{
	Use:   "ec2 init",
	Short: "Initialize for EC2 price list",
	RunE:  doEC2Init,
}

func doEC2Init(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	ec2Cmd.AddCommand(ec2InitCmd)
}
