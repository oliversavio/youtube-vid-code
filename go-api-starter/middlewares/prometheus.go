package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterPrometheusMetrics() {
	prometheus.MustRegister(latency)
}

func RecordRequestLatency(c *fiber.Ctx) error {
	start := time.Now()
	next := c.Next()
	elapsed := time.Since(start).Seconds()

	latency.WithLabelValues(
		c.Route().Method,
		c.Route().Path,
	).Observe(elapsed)

	return next
}

var latency = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Namespace:  "api",
		Name:       "latency_seconds",
		Help:       "Latency distributions.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"method", "path"},
)
