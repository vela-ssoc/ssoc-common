package stegano

import (
	"archive/zip"
	"context"
	"encoding/json"
	"io"
	"io/fs"
)

func AddFS(w io.Writer, fsys fs.FS, offset int64) error {
	zw := zip.NewWriter(w)
	defer zw.Close()
	if offset > 0 {
		zw.SetOffset(offset)
	}

	return zw.AddFS(fsys)
}

func Open(f string) (*zip.ReadCloser, error) {
	return zip.OpenReader(f)
}

func ReadManifest(f string, v any) error {
	zrc, err := Open(f)
	if err != nil {
		return err
	}
	defer zrc.Close()

	// manifest.json 为约定的隐写配置文件名字，不要随意改变。
	mf, err := zrc.Open("manifest.json")
	if err != nil {
		return err
	}
	defer mf.Close()

	return json.NewDecoder(mf).Decode(v)
}

type File[T any] string

func (f File[T]) Read(context.Context) (*T, error) {
	t := new(T)
	if err := ReadManifest(string(f), t); err != nil {
		return nil, err
	}

	return t, nil
}
