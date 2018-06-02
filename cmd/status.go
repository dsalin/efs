package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display current status of working group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("efsctl: status")
	},
}
