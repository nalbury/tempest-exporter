package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature",
		Help: "Temperature in Celcius",
	})
	prometheus.MustRegister(temp)
	temp.Set(69.0)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:6969", nil)
}
