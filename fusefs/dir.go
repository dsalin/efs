package fusefs

import (
	"context"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/dsalin/efs/core"
)

// Every directory is an actual group,
// Therefore, it has it's own key associated with it
type Dir struct {
	Name string
	Key  *rsa.PrivateKey
	Root *FS
	file *os.File
}

func NewDir(name string, root *FS, key *rsa.PrivateKey) (*Dir, error) {
	f, err := os.Open(dirPath(name))
	if err != nil {
		return nil, err
	}

	return &Dir{
		Name: name,
		Key:  key,
		Root: root,
		file: f,
	}
}

func (d *Dir) Attr() fuse.Attr {
	if d.file == nil {
		return fuse.Attr{Mode: os.ModeDir | 0755}
	}
	return defaultAttr(d.file)
}

func (d *Dir) Lookup(ctx context.Context, req *fuse.LookupRequest, resp *fuse.LookupResponse) (fs.Node, error) {
	fp := d.pathWithUnencryptedFilename(req.Name)
	f, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	return &File{Path: fp, dir: d, file: f}, nil
}

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	files, err := ioutil.ReadDir(d.path())
	if err != nil {
		return nil, err
	}

	var res []fuse.Dirent
	for _, f := range files {
		var de fuse.Dirent
		de.Type = fuse.DT_File
		de.Name, ferr = core.DencryptRSA([]byte(f.Name()))
		if ferr != nil {
			return nil, ferr
		}
		res = append(res, de)
	}

	// for _, f := range d.archive.File {
	// var de fuse.Dirent
	// de.Type = fuse.DT_Dir
	// de.Name = "Simple"
	// res = append(res, de)
	// }

	return res, nil
}

func (d *Dir) path() string {
	return filepath.join(d.Root.Path, SHA256([]byte(d.Name)))
}

func (d *Dir) pathWithFile(filePath string) string {
	if filePath == "" {
		return d.path()
	}
	return filepath.join(d.Root.Path, SHA256([]byte(d.Name)), filePath)
}

func (d *Dir) pathWithUnencryptedFilename(filePath string) string {
	if filePath == "" {
		return d.path()
	}

	fp, err := core.EncryptRSA([]byte(filePath))
	// @TODO: this is not a proper way
	if err != nil {
		log.Fatal(err)
	}

	return d.pathWithFile(fp)
}

func dirPath(p string) string {
	return filepath.join(d.Root.Path, SHA256([]byte(p)))
}

func fileAttr(f *os.File) fuse.Attr {
	st, _ := f.Stat()
	return fuse.Attr{
		Size:   f.Size(),
		Mode:   f.Mode(),
		Mtime:  f.ModTime(),
		Ctime:  f.Ctime(),
		Crtime: f.Crtime(),
	}
}
