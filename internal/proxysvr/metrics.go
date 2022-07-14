package proxysvr

import "github.com/prometheus/client_golang/prometheus"

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "metaetcd_request_count",
			Help: "Number of total requests partitioned by method.",
		},
		[]string{"method"},
	)

	getMemberRevDepth = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "metaetcd_get_member_rev_depth",
			Help: "Depth of recursion when mapping meta cluster revision to a specific member cluster.",
		})
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(getMemberRevDepth)
}
