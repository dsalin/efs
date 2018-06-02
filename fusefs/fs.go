package fusefs

import (
	"crypto/rsa"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type FS struct {
	// directory in which the file system is initialized
	RootDir *os.File
	// Default key to access the root directory
	AccessKey *rsa.PrivateKey
}

func (f *FS) Root() (fs.Node, fuse.Error) {
	return &Dir{
		Name: "default",
		Key:  f.AccessKey,
		Root: f,
	}, nil
}
