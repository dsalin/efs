package cmd

import (
	"fmt"
	"os"

	"github.com/dsalin/efs/core"
	"github.com/spf13/cobra"
)

const permissionBits os.FileMode = 0755

// path to initialize the EFS
var source, mainAccessKey string

func init() {
	initCmd.Flags().StringVarP(&source, "source", "s", "", "Source device/file path to initialize EFS into")
	initCmd.Flags().StringVarP(&mainAccessKey, "key", "k", "", "Main Access Key path to initialize EFS into")
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize EFS on path",
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" {
			fmt.Printf("efsctl panic: no source path supplied\n")
			return
		}
		if mainAccessKey == "" {
			fmt.Printf("efsctl panic: no main access key path supplied\n")
			return
		}
		fmt.Printf("Source Device/File Path: %s\nMain Access Key Path: %s\n", source, mainAccessKey)

		os.MkdirAll("./.efs", permissionBits)
		def := createFileOrPanic("./.efs/DEFAULT")
		defer def.Close()
		head := createFileOrPanic("./.efs/HEAD")
		defer head.Close()
		conf := createFileOrPanic("./.efs/config")
		defer conf.Close()

		privateKey, sha, err := core.ReadPrivateKey(mainAccessKey)
		if err != nil {
			fmt.Printf("efsctl panic: %s\n", err)
			return
		}
		fmt.Printf("Private Key SHA: %s\n", sha)
		_, derr := def.WriteString(sha + "\n")
		if derr != nil {
			fmt.Printf("efsctl panic: %s\n", derr)
			return
		}

		encrypted, eerr := core.EncryptRSA([]byte("HELLO WORLD!"), privateKey)
		if eerr != nil {
			fmt.Printf("efsctl panic: %s\n", eerr)
			return
		}
		fmt.Println("efs: %v\n", encrypted)

		decrypted, derr := core.DencryptRSA(encrypted, privateKey)
		if eerr != nil {
			fmt.Printf("efsctl panic: %s\n", derr)
			return
		}
		fmt.Println("efs: %v\n", string(decrypted))

		fmt.Println("efsctl: directory initialized")
	},
}

func createFileOrPanic(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}
