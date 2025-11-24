package stegano_test

import (
	"io"
	"os"
	"testing"

	"github.com/vela-ssoc/ssoc-common/stegano"
)

func TestAddFS(t *testing.T) {
	const baseFile = "base.png"
	const outFile = "out.png"
	const fsDir = "testdata"

	out, err := os.Create(outFile)
	if err != nil {
		t.Error(err)
		return
	}
	defer out.Close()

	base, err := os.Open(baseFile)
	if err != nil {
		t.Error(err)
		return
	}
	defer base.Close()

	fsys := os.DirFS(fsDir)
	offset, err := io.Copy(out, base)
	if err != nil {
		t.Error(err)
		return
	}

	if err = stegano.AddFS(out, fsys, offset); err != nil {
		t.Error(err)
		return
	}

	t.Log("OK")
}
