package main

import (
	"crypto/rsa"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dsalin/efs/fusefs"
)

// We assume the zip file contains entries for directories too.
var progName = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", progName)
	fmt.Fprintf(os.Stderr, "  %s ZIP MOUNTPOINT\n", progName)
	flag.PrintDefaults()
}

func Init(path, mountpoint string, akey *rsa.PrivateKey) error {
	if path == "" || mountpoint == "" {
		return fmt.Errorf("[FuseFS Init] Path or mountpoint cannot be empty")
	}

	if err := fusefs.Mount(path, mountpoint); err != nil {
		log.Fatal(err)
	}
}
