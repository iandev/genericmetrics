Add a new type to metrics/types.go
```go
type FooMetric struct {
	Tag1 string
	Tag2 string
}
```

Add it to the MetricTypes struct
```go
type MetricTypes struct {
	BarBazBing BarBazBing
	Baz        Baz
	BazBar     BazBar
	BarBing    BarBing
	FooMetric  FooMetric `type:"counter" prometheus:"help=help message"`
}
```

run:
```bash
go generate ./...
```

this generates new type and functions in metrics/funcs.go
```go
type FooMetricCounter struct {
	FooMetric *FooMetric
	c prometheus.Counter
}
// NewFooMetric returns an instance of a FooMetricCounter and registers the counter with prometheus
func NewFooMetricCounter(m *FooMetric) FooMetricCounter {
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
	return FooMetricCounter{
		FooMetric: m,
		c: counter,
	}
}
// Inc is a wrapper around the prometheus Inc() method
func (m FooMetricCounter) Inc() {
	m.c.Inc()
}
```

use the new metric:
```go
foo := metrics.FooMetric{
	Tag1: "bar",
	Tag2: "baz",
}
counter := metrics.NewFooMetricCounter(&foo)
counter.Inc()
```

Run sample:
```bash
make run
```
