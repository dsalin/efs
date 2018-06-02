package fusefs

import (
	"crypto/rsa"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func Mount(path, mountpoint string, akey *rsa.PrivateKey) error {
	rootDir, ferr := os.Open(path)
	if ferr != nil {
		return fmr.Errorf("[FuseFS Mount] Path error %v\n", ferr)
	}

	if akey == nil {
		return fmr.Errorf("[FuseFS Mount] Access Key not provided")
	}

	c, err := fuse.Mount(mountpoint)
	if err != nil {
		return err
	}
	defer c.Close()

	filesys := &FS{
		RootDir:   rootDir,
		AccessKey: akey,
	}
	if err := fs.Serve(c, filesys); err != nil {
		return err
	}

	// check if the mount process has an error to report
	<-c.Ready
	if err := c.MountError; err != nil {
		return err
	}

	return nil
}
