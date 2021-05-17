package main

import (
	conf "github.com/aserdean/gitlab-exporter/config"
	"github.com/aserdean/gitlab-exporter/exporter"
	"github.com/aserdean/gitlab-exporter/http"
	"github.com/fatih/structs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	log            *logrus.Logger
	applicationCfg conf.Config
	mets           map[string]*prometheus.Desc
)

func init() {
	applicationCfg = conf.Init()
	mets = exporter.AddMetrics()
	log = logrus.New()

}

func main() {
	log.WithFields(structs.Map(applicationCfg)).Info("Starting Exporter")

	exp := exporter.Exporter{
		APIMetrics: mets,
		Config:     applicationCfg,
	}

	http.NewServer(exp).Start()
}
