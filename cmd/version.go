package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "efsctl: terminal based tool to interact with EFS",
	Long:  `EFS is an encrypted Linux filesystem. Primary difference of EFS is that all files are encrypted group based, not file based.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("efsctl: version 1.0.0 (dev)")
	},
}
