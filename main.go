package main

import (
	"github.com/iandev/genericmetrics/metrics"
)

func main() {
	BarBazBing := metrics.BarBazBing{
		Bar:  "assd",
		Baz:  "dsasd",
		Bing: "dsasd",
	}
	counter := metrics.NewBarBazBingCounter(&BarBazBing)
	counter.Inc()
}
