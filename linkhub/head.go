package linkhub

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

func ReadAuth(r io.Reader, v any) error {
	head := make([]byte, 4)
	n, err := io.ReadFull(r, head)
	if err != nil {
		return err
	} else if n != 4 {
		return io.ErrShortWrite
	}

	size := binary.BigEndian.Uint32(head)
	data := make([]byte, size)
	if n, err = io.ReadFull(r, data); err != nil {
		return err
	} else if n != int(size) {
		return io.ErrShortWrite
	}

	return json.Unmarshal(data, v)
}

func WriteAuth(w io.Writer, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	n := len(data)
	if n > 65535 {
		return io.ErrShortWrite
	}
	head := make([]byte, 4)
	binary.BigEndian.PutUint32(head, uint32(n))
	if _, err = w.Write(head); err != nil {
		return err
	}
	_, err = w.Write(data)

	return err
}
