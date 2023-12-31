package vrctFs

import "io/fs"

type DirEntry struct {
	name   string
	isDir  bool
	type_  fs.FileMode
	StatFn func() (fs.FileInfo, error)
}

func (d DirEntry) Name() string {
	return d.name
}

func (d DirEntry) IsDir() bool {
	return d.isDir
}

func (d DirEntry) Type() fs.FileMode {
	return d.type_
}

func (d DirEntry) Info() (fs.FileInfo, error) {
	return d.StatFn()
}
