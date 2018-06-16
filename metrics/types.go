package metrics

//go:generate go run ../cmd/generate/main.go

type MetricTypes struct {
	BarBazBing BarBazBing `prometheus:"help=HELPME"`
	Baz        Baz        `prometheus:"help=Baz HELP"`
	BazBar     BazBar     `prometheus:"help=BazBar HELP"`
	BarBing    BarBing    `prometheus:"help=BarBing HELP"`
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
