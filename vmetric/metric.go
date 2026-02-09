package vmetric

import (
	"context"
	"io"
	"os"
	"runtime"
	"strconv"

	"github.com/VictoriaMetrics/metrics"
)

type MetricWriter interface {
	WriteMetric(ctx context.Context, w io.Writer)
}

type ConfigLoader interface {
	LoadConfig(ctx context.Context) (pushURL string, opts *metrics.PushOptions, err error)
}

func ManagerLabel() string {
	hostname, _ := os.Hostname()
	return "instance=" + strconv.Quote("manager") +
		",instance_type=" + strconv.Quote("manager") +
		",instance_name=" + strconv.Quote(hostname) +
		",goos=" + strconv.Quote(runtime.GOOS) +
		",goarch=" + strconv.Quote(runtime.GOARCH)
}

func BrokerLabel(id, name string) string {
	return "instance=" + strconv.Quote(id) +
		",instance_type=" + strconv.Quote("broker") +
		",instance_name=" + strconv.Quote(name) +
		",goos=" + strconv.Quote(runtime.GOOS) +
		",goarch=" + strconv.Quote(runtime.GOARCH)
}
