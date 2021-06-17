package main

import (
	sonar "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/sonar"
	config "bitbucket.org/ciandt_it/metricas-alta-renda/pkg/config"
)

func main() {
	config.InitConfig()
	sonar.CollectMetric()
}
