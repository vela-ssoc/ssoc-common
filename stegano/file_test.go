package stegano_test

import (
	"io"
	"io/fs"
	"os"
	"testing"

	"github.com/vela-ssoc/ssoc-common/stegano"
)

func TestRead(t *testing.T) {
	zfs, err := stegano.Open("outfile.exe")
	if err != nil {
		t.Error(err)
		return
	}
	defer zfs.Close()

	fs.WalkDir(zfs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		t.Log(path, d.Name())

		return nil
	})
}

// steganoed

func TestAppend(t *testing.T) {
	base, err := os.Open("example.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer base.Close()

	out, err := os.Create("outfile.exe")
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, base); err != nil {
		t.Error(err)
		return
	}

	fsys, err := os.OpenRoot("testdata")
	if err != nil {
		t.Error(err)
		return
	}
	defer fsys.Close()

	if err = stegano.Append(out, fsys.FS()); err != nil {
		t.Error(err)
		return
	}
}

func TestZip(t *testing.T) {
	// _ = zw.SetOffet(filesize)
}
