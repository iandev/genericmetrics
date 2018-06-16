package metrics

import "github.com/prometheus/client_golang/prometheus"

// BarBazBingCounterVec gets a prometheus Counter for metric BarBazBing
func BarBazBingCounterVec(BarBazBing BarBazBing) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "BarBazBing", Help: ""}, []string{"Bar", "Baz", "Bing"}).
		WithLabelValues(BarBazBing.Bar, BarBazBing.Baz, BarBazBing.Bing)
}


// BazCounterVec gets a prometheus Counter for metric Baz
func BazCounterVec(Baz Baz) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "Baz", Help: ""}, []string{"Baz"}).
		WithLabelValues(Baz.Baz)
}


// BazBarCounterVec gets a prometheus Counter for metric BazBar
func BazBarCounterVec(BazBar BazBar) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "BazBar", Help: ""}, []string{"Baz", "Bar"}).
		WithLabelValues(BazBar.Baz, BazBar.Bar)
}


// BarBingCounterVec gets a prometheus Counter for metric BarBing
func BarBingCounterVec(BarBing BarBing) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "BarBing", Help: ""}, []string{"Bar", "Bing"}).
		WithLabelValues(BarBing.Bar, BarBing.Bing)
}

