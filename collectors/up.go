package collectors

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

type UpCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewUpCollector(db *sql.DB) *UpCollector {
	prometheus.NewD
}

func (c *UpCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc

}

func (c *UpCollector) Collect(metrics chan<- prometheus.Metrics) {
	up := 1
	if err := c.db.Ping(); err != nil {
		up = 0
	}
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, float64(up))
}
