package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/bounceexchange/genericmetrics/metrics"
)

var funcTemplate = `
// %[1]sCounterVec gets a prometheus Counter for metric %[1]s
func %[1]sCounterVec(%[1]s %[1]s) prometheus.Counter {
	return prometheus.
		NewCounterVec(prometheus.CounterOpts{Name: "%[1]s", Help: ""}, []string{%[2]s}).
		WithLabelValues(%[3]s)
}
`

func main() {
	out, err := os.Create("metrics/funcs.go")
	if err != nil {
		log.Fatalf("Cannot create funcs.go file")
	}

	defer out.Close()
	out.WriteString("package metrics\n\n")
	out.WriteString(`import "github.com/prometheus/client_golang/prometheus"`)
	out.WriteString("\n")

	metrics := metrics.MetricTypes{}

	e := reflect.ValueOf(&metrics).Elem()

	for i := 0; i < e.NumField(); i++ {
		varType := e.Type().Field(i).Type
		metricName := varType.Name()
		//typeName := varType.String()

		labels := make([]string, varType.NumField())
		values := make([]string, varType.NumField())

		for j := 0; j < varType.NumField(); j++ {
			fieldName := varType.Field(j).Name
			labels[j] = `"` + fieldName + `"`
			values[j] = metricName + "." + fieldName
		}

		fn := fmt.Sprintf(funcTemplate, metricName, strings.Join(labels, ", "), strings.Join(values, ", "))
		out.WriteString(fn)
		out.WriteString("\n")
	}

	out.Sync()
}
