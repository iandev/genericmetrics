package metrics

//go:generate go run ../cmd/generate/main.go -p prometheus -o funcs.go

type MetricTypes struct {
	BarBazBing BarBazBing `methods:"inc, desc" prometheus:"help=HELPME, foo=bar"`
	Baz        Baz        `methods:"inc, desc" prometheus:"help=Baz HELP"`
	BazBar     BazBar     `methods:"inc, desc" prometheus:"help=BazBar HELP"`
	BarBing    BarBing    `methods:"inc, desc" prometheus:"help=BarBing HELP"`
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
