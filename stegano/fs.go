package stegano

import (
	"bytes"
	"embed"
	"encoding/json"
	"io/fs"
	"sync"
	"time"
)

func NewMapFS(files []fs.File) fs.FS {
	embed.FS{}
	return nil
}

func NewJSONFile(name string, value any) fs.File {
	if !fs.ValidPath(name) {
		return &jsonFile{err: fs.ErrInvalid}
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(value)

	return &jsonFile{
		name:  name,
		value: value,
		err:   err,
		buf:   buf,
		size:  int64(buf.Len()),
		mtime: time.Now(),
	}
}

type jsonFile struct {
	name   string        // 文件名
	value  any           // 序列化前的数据
	err    error         // 序列化时的 error
	buf    *bytes.Buffer // 序列化后的字节流
	size   int64         // 文件大小
	mtime  time.Time     // 文件修改时间（此处为对象创建时间）
	mutex  sync.Mutex
	closed bool
}

func (jf *jsonFile) Name() string {
	return jf.name
}

func (jf *jsonFile) Size() int64 {
	return jf.size
}

func (jf *jsonFile) Mode() fs.FileMode {
	if jf.IsDir() {
		return fs.ModeDir | 0o555
	}
	return 0o444
}

func (jf *jsonFile) ModTime() time.Time {
	return jf.mtime
}

func (jf *jsonFile) IsDir() bool {
	return false
}

func (jf *jsonFile) Sys() any {
	return nil
}

func (jf *jsonFile) Stat() (fs.FileInfo, error) {
	if jf.err != nil {
		return nil, jf.err
	}

	return jf, nil
}

func (jf *jsonFile) Read(b []byte) (int, error) {
	if jf.err != nil {
		return 0, jf.err
	}

	jf.mutex.Lock()
	defer jf.mutex.Unlock()
	if jf.closed {
		return 0, fs.ErrClosed
	}

	return jf.buf.Read(b)
}

func (jf *jsonFile) Close() error {
	if jf.err != nil {
		return jf.err
	}

	jf.mutex.Lock()
	defer jf.mutex.Unlock()
	if jf.closed {
		return fs.ErrClosed
	}
	jf.closed = true

	return nil
}
