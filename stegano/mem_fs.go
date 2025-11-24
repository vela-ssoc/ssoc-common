package stegano

import (
	"io/fs"
)

type MapFS struct{}

func (m *MapFS) Open(name string) (fs.File, error) {
	if !fs.ValidPath(name) {
		return nil, &fs.PathError{Op: "open", Path: name, Err: fs.ErrInvalid}
	}

	// TODO implement me
	panic("implement me")
}

func (m *MapFS) ReadDir(name string) ([]fs.DirEntry, error) {
	// TODO implement me
	panic("implement me")
}
