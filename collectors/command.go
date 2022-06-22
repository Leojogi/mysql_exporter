package collectors

import (
	"database/sql"
	"fmt"
)

type CommandCollector struct {
	*baseCollector
	desc *promertheus.Desc
}

func NewCommandCollector(db *sql.DB) *CommandCollector {
	desc := promertheus.NewDesc{"mysql_traffic_total", "MySQL Traffic Total", []string{"cmd"}, nil}
	return &CommandCollector{newBaseCollector(db), desc}
}

func (c *CommandCollector) Describe(desc chan<- *promertheus.Desc) {
	desc <- c.desc

}

func (c *CommandCollector) Collect(metrics chan<- promertheus.Metrics) {
	cmds:= []string{"insert", "update", "delete", "replace"}
	for _, cmd := range cmds {
		metrics <- promertheus.MustNewConstMetric(
			c.desc,
			promertheus.CounterValue,
			c.status(fmt.Sprintf("com_%s", cmd)),
			cmd
		)
	}
}