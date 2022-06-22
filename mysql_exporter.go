package main

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {

	addr := ":8999"
	mysqlAddr := "localhost:3306"
	dsn := "golang:golang@2022@tcp(localhost:3306)/mysql?charset=utf8mb4&loc=PRC&parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Fatel(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	mysqlInfo := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "mysql_info",
		Help:        "mysql info",
		ConstLables: prometheus.Lables{"addr": mysqlAddr},
	})
	mysqlInfo.Set(1)

	//定义指标

	//注册指标
	//1.时间出发 2.业务请求触发 3.metrics请求触发
	prometheus.MustRegister(collectors.NewUpCollector(db))
	prometheus.MustRegister(collectors.NewSlowQueriesCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
	prometheus.MustRegister(collectors.NewConnectionCollector(db))
	prometheus.MustRegister(collectors.CommandCollector(db))
	prometheus.MustRegister(mysqlInfo)

	//注册控制器
	http.Handle("/metrics", promhttp.Handler())

	//启动web服务
	http.ListenAndServe(addr, nil)

}
