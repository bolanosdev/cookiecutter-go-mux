package prometheus

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils"

	"github.com/IBM/pgxpoolprometheus"
	"github.com/gorilla/mux"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	metricCPUUserTotal    = "cpu_user_total"
	metricCPUSystemTotal  = "cpu_system_total"
	metricCPUIdleTotal    = "cpu_idle_total"
	metricMemUsedTotal    = "mem_used_total"
	metricMemCachedTotal  = "mem_cached_total"
	metricRequestTotal    = "request_total"
	metricRequestUV       = "request_uv"
	metricRequestUVTotal  = "request_uv_total"
	metricURIRequestTotal = "uri_request_total"
	metricRequestBody     = "request_body_total"
	metricResponseBody    = "response_body_total"
	metricRequestDuration = "request_duration"
	metricSlowRequest     = "slow_request_total"

	bloomFilter *utils.BloomFilter
)

func (m *Monitor) Interceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.RequestURI
		startTime := time.Now()

		if url == m.metricPath {
			m.hostMetrics()
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
		m.metricHandle(w, r, startTime)
	})
}

// Use set gin metrics middleware
func (m *Monitor) Use(r *mux.Router) {
	m.initMetrics()
	collector := pgxpoolprometheus.NewCollector(m.pool, map[string]string{"db_name": "brawney_db"})
	prometheus.MustRegister(collector)
	r.Use(m.Interceptor)
	r.Handle(m.metricPath, promhttp.Handler())
}

// initMetrics used to init gin metrics
func (m *Monitor) initMetrics() {
	bloomFilter = utils.NewBloomFilter()

	_ = monitor.AddMetric(&Metric{
		Type:        Histogram,
		Name:        metricCPUUserTotal,
		Description: "Total CPU time consumed by the user process.",
		Labels:      []string{},
		Buckets:     m.cpuUsage,
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Histogram,
		Name:        metricCPUSystemTotal,
		Description: "Total CPU time consumed by the system process.",
		Labels:      []string{},
		Buckets:     m.cpuUsage,
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Histogram,
		Name:        metricCPUIdleTotal,
		Description: "Total CPU idle process.",
		Labels:      []string{},
		Buckets:     m.cpuUsage,
	})

	_ = monitor.AddMetric(&Metric{
		Type:        Histogram,
		Name:        metricMemUsedTotal,
		Description: "Used memory",
		Labels:      []string{},
		Buckets:     m.memUsage,
	})

	_ = monitor.AddMetric(&Metric{
		Type:        Histogram,
		Name:        metricMemCachedTotal,
		Description: "Cached memory",
		Labels:      []string{},
		Buckets:     m.memUsage,
	})

	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricRequestTotal,
		Description: "all the server received request num.",
		Labels:      m.getMetricLabelsIncludingMetadata(metricRequestTotal),
	})

	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricRequestUV,
		Description: "all the server received ip num.",
		Labels:      m.getMetricLabelsIncludingMetadata(metricRequestUV),
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricRequestUVTotal,
		Description: "all the server received ip num.",
		Labels:      m.getMetricLabelsIncludingMetadata(metricRequestUVTotal),
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricURIRequestTotal,
		Description: "all the server received request num with every uri.",
		Labels:      m.getMetricLabelsIncludingMetadata(metricURIRequestTotal),
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricRequestBody,
		Description: "the server received request body size, unit byte",
		Labels:      m.getMetricLabelsIncludingMetadata(metricRequestBody),
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricResponseBody,
		Description: "the server send response body size, unit byte",
		Labels:      m.getMetricLabelsIncludingMetadata(metricResponseBody),
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Histogram,
		Name:        metricRequestDuration,
		Description: "the time server took to handle the request.",
		Labels:      m.getMetricLabelsIncludingMetadata(metricRequestDuration),
		Buckets:     m.reqDuration,
	})
	_ = monitor.AddMetric(&Metric{
		Type:        Counter,
		Name:        metricSlowRequest,
		Description: fmt.Sprintf("the server handled slow requests counter, t=%d.", m.slowTime),
		Labels:      m.getMetricLabelsIncludingMetadata(metricSlowRequest),
	})
}

func (m *Monitor) includesMetadata() bool {
	return len(m.metadata) > 0
}

func (m *Monitor) getMetadata() ([]string, []string) {
	metadata_labels := []string{}
	metadata_values := []string{}

	for v := range m.metadata {
		metadata_labels = append(metadata_labels, v)
		metadata_values = append(metadata_values, m.metadata[v])
	}

	return metadata_labels, metadata_values
}

func (m *Monitor) getMetricLabelsIncludingMetadata(metricName string) []string {
	includes_metadata := m.includesMetadata()
	metadata_labels, _ := m.getMetadata()

	switch metricName {
	case metricRequestDuration:
		metric_labels := []string{"uri"}
		if includes_metadata {
			metric_labels = append(metric_labels, metadata_labels...)
		}
		return metric_labels

	case metricRequestUV:
		metric_labels := []string{"clientIP"}
		if includes_metadata {
			metric_labels = append(metric_labels, metadata_labels...)
		}
		return metric_labels

	case metricURIRequestTotal:
		metric_labels := []string{"uri", "method", "code"}
		if includes_metadata {
			metric_labels = append(metric_labels, metadata_labels...)
		}
		return metric_labels

	case metricSlowRequest:
		metric_labels := []string{"uri", "method", "code"}
		if includes_metadata {
			metric_labels = append(metric_labels, metadata_labels...)
		}
		return metric_labels

	default:
		var metric_labels []string = nil
		if includes_metadata {
			metric_labels = metadata_labels
		}
		return metric_labels
	}
}

func (m *Monitor) hostMetrics() {
	cpu, err := cpu.Get()
	if err != nil {
		log.Print(err)
		return
	}
	memory, err := memory.Get()
	if err != nil {
		log.Print(err)
		return
	}

	cpu_user := (float64(cpu.User) / float64(cpu.Total)) * 100
	cpu_system := (float64(cpu.System) / float64(cpu.Total)) * 100
	cpu_idle := (float64(cpu.Idle) / float64(cpu.Total)) * 100

	// EmitMemoryMetrics(c.packageName, memory.Total, memory.Used, memory.Cached, memory.Free)
	memory_used := (float64(memory.Used) / float64(memory.Total) * 100)
	memory_cached := (float64(memory.Cached) / float64(memory.Total) * 100)

	var metric_values []string = nil
	_ = m.GetMetric(metricCPUUserTotal).Observe(m.getMetricValues(metric_values), cpu_user)
	_ = m.GetMetric(metricCPUSystemTotal).Observe(m.getMetricValues(metric_values), cpu_system)
	_ = m.GetMetric(metricCPUIdleTotal).Observe(m.getMetricValues(metric_values), cpu_idle)
	_ = m.GetMetric(metricCPUUserTotal).Observe(m.getMetricValues(metric_values), cpu_user)
	_ = m.GetMetric(metricMemUsedTotal).Observe(m.getMetricValues(metric_values), memory_used)
	_ = m.GetMetric(metricMemCachedTotal).Observe(m.getMetricValues(metric_values), memory_cached)

	// log.Printf("cpu user %v system %v, idle %v", cpu_user, cpu_system, cpu_idle)
	// log.Printf("mem used %v cached %v, ", memory_used, memory_cached)
}

func (m *Monitor) metricHandle(w http.ResponseWriter, r *http.Request, start time.Time) {
	status := w.Header().Get("x-status-code")

	// set request total
	var metric_values []string = nil
	_ = m.GetMetric(metricRequestTotal).Inc(m.getMetricValues(metric_values))

	// set uv
	if clientIP := r.RemoteAddr; !bloomFilter.Contains(clientIP) {
		clientIP = strings.Split(clientIP, ":")[0]
		bloomFilter.Add(clientIP)
		metric_values = nil
		_ = m.GetMetric(metricRequestUVTotal).Inc(m.getMetricValues(metric_values))

		metric_values = []string{clientIP}
		_ = m.GetMetric(metricRequestUV).Inc(m.getMetricValues(metric_values))
	}

	// set uri request total
	metric_values = []string{r.RequestURI, r.Method, status}
	_ = m.GetMetric(metricURIRequestTotal).Inc(m.getMetricValues(metric_values))

	// set request body size
	// since r.ContentLength can be negative (in some occasions) guard the operation
	if r.ContentLength >= 0 {
		metric_values = nil
		_ = m.GetMetric(metricRequestBody).Add(m.getMetricValues(metric_values), float64(r.ContentLength))
	}

	// set slow request
	latency := time.Since(start)
	if int32(latency.Seconds()) > m.slowTime {
		metric_values = []string{r.RequestURI, r.Method, status}
		_ = m.GetMetric(metricSlowRequest).Inc(m.getMetricValues(metric_values))
	}

	// set request duration
	metric_values = []string{r.RequestURI}
	_ = m.GetMetric(metricRequestDuration).Observe(m.getMetricValues(metric_values), latency.Seconds())

	//// set response size
	//if w.Size() > 0 {
	//metric_values = nil
	//_ = m.GetMetric(metricResponseBody).Add(m.getMetricValues(metric_values), float64(w.Size()))
	//}
}

func (m *Monitor) getMetricValues(metric_values []string) []string {
	includes_metadata := m.includesMetadata()
	_, metadata_values := m.getMetadata()
	if includes_metadata {
		metric_values = append(metric_values, metadata_values...)
	}
	return metric_values
}
