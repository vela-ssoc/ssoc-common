package stegano_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"testing"

	"github.com/vela-ssoc/ssoc-common/stegano"
)

type manifest struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Random string `json:"random"`
}

// TestAddManifest 向二进制中隐写数据。
func TestAddManifest(t *testing.T) {
	const baseFile = "base.png"
	const outFile = "out.png"

	buf := make([]byte, 10)
	_, _ = rand.Read(buf)
	random := hex.EncodeToString(buf)
	t.Logf("生成随机密钥：%s", random)
	payload := &manifest{
		Name:   "Alice",
		Age:    24,
		Random: random,
	}

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

	offset, err := io.Copy(out, base)
	if err != nil {
		t.Error(err)
		return
	}

	if err = stegano.AddManifest(out, payload, offset); err != nil {
		t.Error(err)
		return
	}

	t.Log("隐写成功")
}

func TestReadManifest(t *testing.T) {
	const outFile = "out.png"

	bin := stegano.Binary[manifest](outFile)
	payload, err := bin.Read(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("读取隐写文件成功：%#v", payload)
	t.Logf("得到随机密钥：%s", payload.Random)
}
