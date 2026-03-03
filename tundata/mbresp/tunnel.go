package mbresp

import (
	"github.com/vela-ssoc/ssoc-proto/muxconn"
)

type TunnelStat struct {
	Name       string                 `json:"name"`
	Module     string                 `json:"module"`
	Cumulative int64                  `json:"cumulative"`
	Active     int64                  `json:"active"`
	RX         uint64                 `json:"rx"`
	TX         uint64                 `json:"tx"`
	Limit      float64                `json:"limit"`
	Unlimit    bool                   `json:"unlimit"`
	Streams    []*muxconn.StreamStats `json:"streams"`
}
