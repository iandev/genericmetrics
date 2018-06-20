### _Dead simple metrics wrapper._ ###



Add a new type to metrics/types.go
```go
type Foo struct {
	Tag1 string
	Tag2 string
}

type Bar struct {
	Tag1 string
}
```

Add it to the MetricTypes struct
```go
type MetricTypes struct {
	Foo Foo `type:"counter" prometheus:"help=help message"`
	Bar Bar `type:"gauge" prometheus:"help=help message"`
}
```

Run
```bash
go generate ./...
```

This generates new type and functions in metrics/funcs.go
```go
type FooCounter struct {
	Foo *Foo
	c prometheus.Counter
}
// NewFooCounter returns an instance of a FooCounter and registers the counter with prometheus
func NewFooCounter(m *Foo) FooCounter {
	labels := []string{
		"Tag1",
		"Tag2",
	}
	opts := prometheus.CounterOpts{Name: "Foo", Help: "help message"}
	counter := prometheus.NewCounterVec(opts, labels).WithLabelValues(
		m.Tag1,
		m.Tag2,
	)
	prometheus.MustRegister(counter)
	return FooCounter{
		Foo: m,
		c: counter,
	}
}
// Inc is a wrapper around the prometheus Inc() method
func (m FooCounter) Inc() {
	m.c.Inc()
}
// Add is a wrapper around the prometheus Add(float64) method
func (m FooCounter) Add(a float64) {
	m.c.Add(a)
}

type BarGauge struct {
	Bar *Bar
	g prometheus.Gauge
}
// NewBarGauge returns an instance of a BarGauge and registers the gauge with prometheus
func NewBarGauge(m *Bar) Gauge {
	labels := []string{
		"Tag1",
	}
	opts := prometheus.GaugeOpts{Name: "Bar", Help: "Bar HELP"}
	guage := prometheus.NewGaugeVec(opts, labels).WithLabelValues(
		m.Tag1,
	)
	prometheus.MustRegister(guage)
	return BarGauge{
		Bar: m,
		g: guage,
	}
}
// Set is a wrapper around the prometheus Set(float64) method
func (g BarGauge) Set(s float64) {
	g.g.Set(s)
}
```

Use the new metric
```go
foo := metrics.Foo{
	Tag1: "bar",
	Tag2: "baz",
}
counter := metrics.NewFooCounter(&foo)
counter.Inc()
counter.Add(3)

bar := metrics.Bar{
	Tag1: "foo",
}
gauge := metrics.NewBarGauge(&bar)
gauge.Set(55.6)
```

Run sample
```bash
make run
```
