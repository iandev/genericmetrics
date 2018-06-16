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
	One           BarBazBing
	Two           Baz
	Three         BazBar
	Four          BarBing
	DoesNotMatter FooMetric
}
```

run:
```bash
go generate
```

creates a new function in metrics/funcs.go
```go
// FooMetricCounter gets a prometheus Counter for metric FooMetric
func FooMetricCounter(FooMetric FooMetric) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "FooMetric", Help: ""}, []string{"Tag1", "Tag2"}).
		WithLabelValues(FooMetric.Tag1, FooMetric.Tag2)
}
```

use the new metric:
```go
m := metrics.FooMetric{
	Tag1: "tag value",
	Tag2: "tag value",
}
counter := metrics.FooMetricCounter(m)
counter.Inc()
```

Run sample:
```bash
make run
```