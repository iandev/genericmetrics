package metrics

//go:generate go run ../cmd/generate/main.go -p prometheus -o funcs.go

type MetricTypes struct {
	BarBazBing BarBazBing `type:"counter" prometheus:"help=HELPME, foo=bar"`
	Baz        Baz        `type:"gauge" prometheus:"help=Baz HELP"`
	BazBar     BazBar     `type:"counter" prometheus:"help=BazBar HELP"`
	BarBing    BarBing    `type:"counter" prometheus:"help=BarBing HELP"`
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
