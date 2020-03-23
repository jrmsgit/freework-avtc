// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package tx

import (
	"github.com/prometheus/client_golang/prometheus"
)

func Hist() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "log_tx",
		Help: "log tx values",
		Buckets: []float64{
			500,
			1000,
			5000,
			9000,
		},
	}, []string{"tx"})
}
