package main

import (
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

func main() {
	log.SetFlags(0)
	log.SetPrefix(progName + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 2 {
		usage()
		os.Exit(2)
	}

	path := flag.Arg(0)
	mountpoint := flag.Arg(1)
	if err := fusefs.Mount(path, mountpoint); err != nil {
		log.Fatal(err)
	}
}
