package model

import (
	"net/http"
	"strings"
)

type MapHeader map[string]string

func (h MapHeader) Canonical() MapHeader {
	ret := make(MapHeader, len(h))
	for k, s := range h {
		k = strings.TrimSpace(k)
		k = http.CanonicalHeaderKey(k)
		ret[k] = s
	}

	return ret
}
