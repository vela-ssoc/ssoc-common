package mbreq

import "golang.org/x/time/rate"

type TunnelLimit struct {
	Unlimit bool    `json:"unlimit"`
	Limit   float64 `json:"limit"`
}

func (l TunnelLimit) Rate() rate.Limit {
	if l.Unlimit {
		return rate.Inf
	}

	return rate.Limit(l.Limit)
}
