package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type SlowQueriesCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewSlowQueriesCollector(db *sql.DB) *SlowQueriesCollector {
	desc := prometheus.NewDesc("mysql_slow_queries_total", "Mysql slow quereties total", nil, nil)
	return &SlowQueriesCollector{newBaseCollector(db), desc}
}

func (c *SlowQueriesCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *SlowQueriesCollector) Collect(metrics chan<- *prometheus.Metrics) {
	count := c.status("slow_queries")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)

}
