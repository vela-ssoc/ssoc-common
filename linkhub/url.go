package linkhub

import "net/url"

const (
	ServerHost       = "server.ssoc.internal"
	BrokerHostSuffix = ".broker.ssoc.internal"
	AgentHostSuffix  = ".agent.ssoc.internal"
)

func NewBrokerToServerURL(path string, ws ...bool) *url.URL {
	return buildURL(ServerHost, path, ws...)
}

func NewBrokerToAgentURL(agentHost, path string, ws ...bool) *url.URL {
	return buildURL(agentHost+AgentHostSuffix, path, ws...)
}

func NewBrokerToAgentIDURL(agentID int64, path string, ws ...bool) *url.URL {
	return NewBrokerToAgentURL(formatID(agentID), path, ws...)
}

func NewServerToBrokerURL(brokerHost, path string, ws ...bool) *url.URL {
	return buildURL(brokerHost+BrokerHostSuffix, path, ws...)
}

func NewServerToBrokerIDURL(brokerID int64, path string, ws ...bool) *url.URL {
	return NewServerToBrokerURL(formatID(brokerID), path, ws...)
}

func buildURL(host, path string, ws ...bool) *url.URL {
	scheme := "http"
	if len(ws) > 0 && ws[0] {
		scheme = "ws"
	}

	return &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
}
