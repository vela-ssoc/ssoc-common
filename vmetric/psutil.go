package vmetric

import (
	"context"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/VictoriaMetrics/metrics"
	pscpu "github.com/shirou/gopsutil/v4/cpu"
	psmem "github.com/shirou/gopsutil/v4/mem"
	psnet "github.com/shirou/gopsutil/v4/net"
)

func NewPsutil() MetricWriter {
	return &psutilMetric{}
}

type psutilMetric struct {
	mutex  sync.Mutex
	lastAt time.Time
	names  map[string]bool
}

func (m *psutilMetric) WriteMetric(ctx context.Context, w io.Writer) {
	if stat, _ := psmem.VirtualMemoryWithContext(ctx); stat != nil {
		metrics.WriteCounterUint64(w, "system_memory_total_bytes", stat.Total)
		metrics.WriteCounterUint64(w, "system_memory_used_bytes", stat.Used)
	}

	if stats, _ := pscpu.PercentWithContext(ctx, 0, false); len(stats) != 0 {
		metrics.WriteGaugeFloat64(w, "system_cpu_usage_percent", stats[0])
	}

	interfaceNames := m.interfaceNames()
	if len(interfaceNames) == 0 {
		return
	}
	if stats, _ := psnet.IOCountersWithContext(ctx, true); stats != nil {
		for _, stat := range stats {
			name := stat.Name
			if _, exists := interfaceNames[name]; !exists {
				continue
			}

			// 网卡名字转义
			name = strconv.Quote(name)
			metrics.WriteCounterUint64(w, "system_network_receive_bytes{interface="+name+"}", stat.BytesRecv)
			metrics.WriteCounterUint64(w, "system_network_transmit_bytes{interface="+name+"}", stat.BytesSent)
		}
	}
}

// interfaceNames 获取网卡名。
//
// 通过缓存避免每次都要查询，过期时间保证网卡名不会永久缓存，防止网卡热插拔/运行中启停。
func (m *psutilMetric) interfaceNames() map[string]bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	now := time.Now()
	if now.Sub(m.lastAt) <= 10*time.Minute {
		return m.names
	}

	names := make(map[string]bool, 4)
	faces, _ := net.Interfaces()
	for _, face := range faces {
		if face.Flags&net.FlagLoopback == 0 &&
			face.Flags&net.FlagUp != 0 &&
			face.Flags&net.FlagRunning != 0 { // 忽略环回网卡/不工作的网卡
			names[face.Name] = true
		}
	}

	m.lastAt = now
	m.names = names

	return names
}
