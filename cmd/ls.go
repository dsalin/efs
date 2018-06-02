package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List contents of current active group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("efs: directory listing")
	},
}
