package fusefs

import (
	"context"
	"io/ioutil"
	"os"

	"bazil.org/fuse"
	"github.com/dsalin/efs/core"
)

type File struct {
	Path string
	// parent directory/group, so that we can extract the key
	dir  *Dir
	file *os.File
}

func (f *File) Attr() fuse.Attr {
	return fileAttr(f.file)
}

func (f *File) Release(ctx context.Context, req *fuse.ReleaseRequest) error {
	return f.file.Close()
}

func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (Handle, error) {
	return f
}

// Need to get current active Key
func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	dat, err := ioutil.ReadFile(f.path)
	if err != nil {
		return nil, err
	}
	// decrypt the file with current directory key
	ddat, derr := core.DencryptRSA(dat, f.Dir.Key)
	if derr != nil {
		return nil, derr
	}

	return ddat, nil
}
