package muxproto

import (
	"net/url"
)

const (
	ManagerHost      = "manager.ssoc.internal"
	BrokerHost       = "broker.ssoc.internal"
	AgentHost        = "agent.ssoc.internal"
	BrokerSuffixHost = "." + BrokerHost
	AgentSuffixHost  = "." + AgentHost
)

func ToManagerURL(path string, ws ...bool) *url.URL {
	return buildURL(ManagerHost, path, ws)
}

func ServerToBrokerURL(brokerID int64, path string, ws ...bool) *url.URL {
	host := resolveHost(brokerID, BrokerSuffixHost)
	return buildURL(host, path, ws)
}

func buildURL(host, path string, ws []bool) *url.URL {
	scheme := "http"
	if len(ws) != 0 && ws[0] {
		scheme = "ws"
	}

	return &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
}
