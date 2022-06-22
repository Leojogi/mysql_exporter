package collectors

import "database/sql"

type ConnectionCollector struct {
	*baseCollector
	maxConnectionDesc    *promertheus.Desc
	threadsConnectedDesc *promertheus.Desc
}

func NewConnectionCollector(db *sql.DB) *ConnectionCollector {
	maxConnectionDesc := promertheus.NewDesc("mysql_global_variables_max_connections", "mysql global variables", nil, nil)
	threadsConnectedDesc := promertheus.NewDesc("mysql_global_status_thread_connected", "mysql global status", nil, nil)
	return &ConnectionCollector{NewConnectionCollector(db), maxConnectionDesc, threadsConnectedDesc}
}

func (c *ConnectionCollector) Describe(descs chan<- *promertheus.Desc) {
	descs <- c..maxConnectionDesc
	descs <- c.threadsConnectedDesc
}

func (c *ConnectionCollector) Collect(metrics chan<- *promertheus.Metrics) {
	metrics <- promertheus.MustNewConstMetric(c.maxConnectionDesc, promertheus.GaugeValue, c.variables("max_connections"))
	metrics <- promertheus.MustNewConstMetric(c.threadsConnectedDesc, promertheus.GaugeValue, c.variables("threads_connected"))
}