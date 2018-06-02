package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "efsctl",
	Short: "efsctl: terminal based tool to interact with EFS",
	Long:  `EFS is an encrypted Linux filesystem. Primary difference of EFS is that all files are encrypted group based, not file based.`,
}
