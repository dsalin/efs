package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var currentGroupKeyPath string

func init() {
	createGroupCmd.Flags().StringVarP(&currentGroupKeyPath, "group_key", "k", "", "EFS Group Key")
	createCmd.AddCommand(createGroupCmd)
	RootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create efs entity: group, etc.",
}

var createGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Create EFS group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" {
			fmt.Printf("efsctl panic: EFS is not initialized\n")
			return
		}
		if currentGroupKeyPath == "" {
			fmt.Printf("efsctl panic: group key is not supplied\n")
			return
		}

		fmt.Printf("efsctl create group: %s\n", args[0])
	},
}
