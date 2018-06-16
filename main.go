package main

import (
	"fmt"

	"github.com/iandev/genericmetrics/metrics"
)

//go:generate go run cmd/generate_metric_funcs/main.go

func main() {
	BarBazBing := metrics.BarBazBing{
		Bar:  "assd",
		Baz:  "dsasd",
		Bing: "dsasd",
	}
	BarBazBingCounter := metrics.BarBazBingCounterVec(BarBazBing)
	BarBazBingCounter.Inc()
	fmt.Printf("foo %v\n", BarBazBingCounter.Desc().String())

	Baz := metrics.Baz{
		Baz: "dsasd",
	}
	BazCounter := metrics.BazCounterVec(Baz)
	BazCounter.Inc()
	fmt.Printf("foo %v\n", BazCounter.Desc().String())

	BazBar := metrics.BazBar{
		Baz: "dsasd",
		Bar: "jbjb",
	}
	BazBarCounter := metrics.BazBarCounterVec(BazBar)
	BazBarCounter.Inc()
	fmt.Printf("foo %v\n", BazBarCounter.Desc().String())

	BarBing := metrics.BarBing{
		Bing: "dsasd",
		Bar:  "jbjb",
	}
	BarBingCounter := metrics.BarBingCounterVec(BarBing)
	BarBingCounter.Inc()
	fmt.Printf("foo %v\n", BarBingCounter.Desc().String())
}
