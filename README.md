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

run:
```bash
go generate ./...
```

this generates new type and functions in metrics/funcs.go
```go
type FooCounter struct {
	Foo *Foo
	c prometheus.Counter
}
// NewFooMetric returns an instance of a FooMetricCounter and registers the counter with prometheus
func NewFooCounter(m *Foo) FooCounter {
	labels := []string{
		"Tag1",
		"Tag2",
	}
	opts := prometheus.CounterOpts{Name: "FooMetric", Help: "help message"}
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
```

use the new metric:
```go
foo := metrics.Foo{
	Tag1: "bar",
	Tag2: "baz",
}
counter := metrics.NewFooCounter(&foo)
counter.Inc()

bar := metrics.Bar{
	Tag1: "foo",
}
gauge := metrics.NewBarGauge(&bar)
gauge.Set(55.6)
```

Run sample:
```bash
make run
```
