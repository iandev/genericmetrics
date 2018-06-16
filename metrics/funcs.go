package metrics

import "github.com/prometheus/client_golang/prometheus"

// Auto generated do not edit

// BarBazBingCounter gets a prometheus Counter for metric BarBazBing
func BarBazBingCounter(BarBazBing BarBazBing) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "BarBazBing", Help: ""}, []string{"Bar", "Baz", "Bing"}).
		WithLabelValues(BarBazBing.Bar, BarBazBing.Baz, BarBazBing.Bing)
}

// BazCounter gets a prometheus Counter for metric Baz
func BazCounter(Baz Baz) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "Baz", Help: ""}, []string{"Baz"}).
		WithLabelValues(Baz.Baz)
}

// BazBarCounter gets a prometheus Counter for metric BazBar
func BazBarCounter(BazBar BazBar) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "BazBar", Help: ""}, []string{"Baz", "Bar"}).
		WithLabelValues(BazBar.Baz, BazBar.Bar)
}

// BarBingCounter gets a prometheus Counter for metric BarBing
func BarBingCounter(BarBing BarBing) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "BarBing", Help: ""}, []string{"Bar", "Bing"}).
		WithLabelValues(BarBing.Bar, BarBing.Bing)
}

