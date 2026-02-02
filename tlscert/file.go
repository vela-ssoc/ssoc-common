package tlscert

import (
	"context"
	"crypto/tls"
)

type File struct {
	certFile, keyFile string
}

func NewFile(certFile, keyFile string) File {
	return File{
		certFile: certFile,
		keyFile:  keyFile,
	}
}

func (f File) Load(context.Context) ([]*tls.Certificate, error) {
	crt, err := tls.LoadX509KeyPair(f.certFile, f.keyFile)
	if err != nil {
		return nil, err
	}

	return []*tls.Certificate{&crt}, nil
}
