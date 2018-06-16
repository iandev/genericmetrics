package metrics

type MetricTypes struct {
	One   BarBazBing
	Two   Baz
	Three BazBar
	Four  BarBing
}

type BarBazBing struct {
	Bar  string
	Baz  string
	Bing string
}

type Baz struct {
	Baz string
}

type BazBar struct {
	Baz string
	Bar string
}

type BarBing struct {
	Bar  string
	Bing string
}
