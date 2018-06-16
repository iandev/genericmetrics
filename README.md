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
	FooMetric  FooMetric `methods:"inc" prometheus:"help=help message"`
}
```

run:
```bash
go generate ./...
```

this generates a new function in metrics/funcs.go
```go
// Inc calls the prometheus Inc function using FooMetric for tags
func (m FooMetric) Inc() {
	labels := []string{
		"Tag1",
		"Tag2",
	}
	opts := prometheus.CounterOpts{Name: "FooMetric", Help: "help message"}
	prometheus.NewCounterVec(opts, labels).WithLabelValues(
		m.Tag1,
		m.Tag2,
	).Inc()
}
```

use the new metric:
```go
foo := metrics.FooMetric{
    Tag1:  "bar",
    Tag2:  "baz",
}
counter := metrics.NewFooMetricCounter(&foo)
counter.Inc()
```

Run sample:
```bash
make run
```
