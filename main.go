package main

import (
	"github.com/iandev/genericmetrics/metrics"
)

func main() {
	barBazBing := metrics.BarBazBing{
		Bar:  "assd",
		Baz:  "dsasd",
		Bing: "dsasd",
	}
	counter := metrics.NewBarBazBingCounter(&barBazBing)
	counter.Inc()

	baz := metrics.Baz{
		Baz: "foobar",
	}
	gauge := metrics.NewBazGauge(&baz)
	gauge.Set(55.6)
}
