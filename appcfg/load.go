package appcfg

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Reader 配置加载器。
type Reader[T any] interface {

	// Read 加载配置文件。
	Read(ctx context.Context) (*T, error)
}

func NewJSON[T any](file string, limit ...int64) Reader[T] {
	var lim int64
	if len(limit) > 0 {
		lim = limit[0]
	}

	return &jsonReader[T]{
		file:  file,
		limit: lim,
	}
}

type jsonReader[T any] struct {
	file  string
	limit int64
}

func (j *jsonReader[T]) Read(context.Context) (*T, error) {
	file := j.file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ret := new(T)
	rd := j.withLimit(f)
	ext := strings.ToLower(filepath.Ext(file))
	if ext != ".jsonc" {
		dec := json.NewDecoder(rd)
		if err = dec.Decode(ret); err != nil {
			return nil, err
		}

		return ret, nil
	}

	raw, err := io.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	dat := translate(raw)
	if err = json.Unmarshal(dat, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (j *jsonReader[T]) withLimit(r io.Reader) io.Reader {
	if j.limit <= 0 {
		return r
	}

	return io.LimitReader(r, j.limit)
}
