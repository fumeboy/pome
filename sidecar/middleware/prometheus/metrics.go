package prometheus

import (
	"context"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/status"
)

// 服务端采样打点
type metrics struct {
	requestCounter *prom.CounterVec
	codeCounter    *prom.CounterVec
	latencySummary *prom.SummaryVec
}

//生成server metrics实例
func NewServerMetrics() *metrics {
	return &metrics{
		requestCounter: promauto.NewCounterVec(
			prom.CounterOpts{
				Name: "server_request_total",
				Help: "Total number of RPCs completed on the server, regardless of success or failure.",
			}, []string{"service", "method"}),
		codeCounter: promauto.NewCounterVec(
			prom.CounterOpts{
				Name: "server_handled_code_total",
				Help: "Total number of RPCs completed on the server, regardless of success or failure.",
			}, []string{"service", "method", "grpc_code"}),
		latencySummary: promauto.NewSummaryVec(
			prom.SummaryOpts{
				Name:       "proc_cost",
				Help:       "RPC latency distributions.",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			},
			[]string{"service", "method"},
		),
	}
}

//生成server metrics实例
func NewClientMetrics() *metrics {
	return &metrics{
		requestCounter: promauto.NewCounterVec(
			prom.CounterOpts{
				Name: "rpc_call_total",
				Help: "Total number of RPCs completed on the server, regardless of success or failure.",
			}, []string{"service", "method"}),
		codeCounter: promauto.NewCounterVec(
			prom.CounterOpts{
				Name: "rpc_code_total",
				Help: "Total number of RPCs completed on the server, regardless of success or failure.",
			}, []string{"service", "method", "grpc_code"}),
		latencySummary: promauto.NewSummaryVec(
			prom.SummaryOpts{
				Name:       "rpc_cost",
				Help:       "RPC latency distributions.",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			},
			[]string{"service", "method"},
		),
	}
}

func (m *metrics) IncrRequest(ctx context.Context, serviceName, methodName string) {
	m.requestCounter.WithLabelValues(serviceName, methodName).Inc()
}

func (m *metrics) IncrCode(ctx context.Context, serviceName, methodName string, err error) {
	st, _ := status.FromError(err)
	m.codeCounter.WithLabelValues(serviceName, methodName, st.Code().String()).Inc()
}

func (m *metrics) Latency(ctx context.Context, serviceName, methodName string, us int64) {

	m.latencySummary.WithLabelValues(serviceName, methodName).Observe(float64(us))
}
