package profile

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Reader[T any] interface {
	Read(ctx context.Context) (*T, error)
}

func NewFile[T any](fp string, limit ...int64) Reader[T] {
	var lim int64
	if len(limit) > 0 && limit[0] > 0 {
		lim = limit[0]
	}

	return &fileReader[T]{
		fp:  fp,
		lim: lim,
	}
}

type fileReader[T any] struct {
	fp  string
	lim int64
}

func (fr fileReader[T]) Read(context.Context) (*T, error) {
	fp := fr.fp
	stat, err := os.Stat(fp)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return fr.readDir(fp)
	}

	return fr.readFile(fp)
}

func (fr fileReader[T]) readFile(fp string) (*T, error) {
	fd, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer fd.Close()

	v := new(T)
	rd := fr.withLimit(fd)
	ext := strings.ToLower(filepath.Ext(fp))
	switch ext {
	case ".json":
		err = json.NewDecoder(rd).Decode(v)
	case ".jsonc":
		raw, err1 := io.ReadAll(rd)
		if err1 != nil {
			return nil, err1
		}
		data := translate(raw)
		err = json.Unmarshal(data, v)
	default:
		err = errors.ErrUnsupported
	}
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (fr fileReader[T]) readDir(dir string) (*T, error) {
	// 按照 .jsonc .json 顺序读取配置文件，直到第一个正确的停止。
	errs := make([]error, 0, 3)
	for _, ext := range []string{".jsonc", ".json"} {
		fp := filepath.Join(dir, "application"+ext)
		if cfg, err := fr.readFile(fp); err == nil {
			return cfg, nil
		} else {
			errs = append(errs, err)
		}
	}

	return nil, errors.Join(errs...)
}

func (fr fileReader[T]) withLimit(r io.Reader) io.Reader {
	if fr.lim > 0 {
		return io.LimitReader(r, fr.lim)
	}

	return r
}
