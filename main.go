package main

import (
	"fmt"

	"github.com/iandev/genericmetrics/metrics"
)

func main() {
	BarBazBing := metrics.BarBazBing{
		Bar:  "assd",
		Baz:  "dsasd",
		Bing: "dsasd",
	}
	BarBazBing.Inc()
	fmt.Printf("BarBazBing: %s\n", BarBazBing.Desc().String())

	Baz := metrics.Baz{
		Baz: "dsasd",
	}
	Baz.Inc()
	fmt.Printf("Baz: %s\n", Baz.Desc().String())

	BazBar := metrics.BazBar{
		Baz: "dsasd",
		Bar: "jbjb",
	}
	BazBar.Inc()
	fmt.Printf("BazBar: %s\n", BazBar.Desc().String())

	BarBing := metrics.BarBing{
		Bing: "dsasd",
		Bar:  "jbjb",
	}
	BarBing.Inc()
	fmt.Printf("BarBing: %s\n", BarBing.Desc().String())
}
