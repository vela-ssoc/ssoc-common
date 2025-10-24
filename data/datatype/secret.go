package datatype

import (
	"log/slog"
	"strings"
)

type Ciphertext string

func (s Ciphertext) LogValue() slog.Value {
	return slog.StringValue(s.String())
}

func (s Ciphertext) String() string {
	n := len(s)
	if n <= 16 {
		return strings.Repeat("*", n)
	}

	str := string(s)

	return str[:4] + "******" + str[n-4:]
}
