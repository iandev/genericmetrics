package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/iandev/genericmetrics/metrics"
)

const Prometheus = "prometheus"
const TypeStructTag = "type"

var prometheusTemplate = `// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// {{ .Timestamp }}
package metrics

import "github.com/prometheus/client_golang/prometheus"

{{- range .Funcs }}
{{if eq .Type "counter"}}
type {{ .MetricName }}Counter struct {
	{{ .MetricName }} *{{ .MetricName }}
	c prometheus.Counter
}
// New{{ .MetricName }}Counter returns an instance of a {{ .MetricName }}Counter and registers the counter with prometheus
func New{{ .MetricName }}Counter(m *{{ .MetricName }}) Counter {
	labels := []string{
	{{- range .MetricTags }}
		"{{ printf "%s" .}}",
	{{- end }}
	}
	opts := prometheus.CounterOpts{Name: "{{ .MetricName }}", Help: "{{ index .Provider "help" }}"}
	counter := prometheus.NewCounterVec(opts, labels).WithLabelValues(
	{{- range .MetricTags }}
		m.{{ printf "%s" .}},
	{{- end }}
	)
	prometheus.MustRegister(counter)
	return {{ .MetricName }}Counter{
		{{ .MetricName }}: m,
		c: counter,
	}
}
// Inc is a wrapper around the prometheus Inc() method
func (c {{ .MetricName }}Counter) Inc() {
	c.c.Inc()
}
// Add is a wrapper around the prometheus Add(float64) method
func (c {{ .MetricName }}Counter) Add(a float64) {
	c.c.Add(a)
}

{{else if eq .Type "gauge"}}

type {{ .MetricName }}Gauge struct {
	{{ .MetricName }} *{{ .MetricName }}
	g prometheus.Gauge
}

// New{{ .MetricName }}Gauge returns an instance of a {{ .MetricName }}Gauge and registers the gauge with prometheus
func New{{ .MetricName }}Gauge(m *{{ .MetricName }}) Gauge {
	labels := []string{
	{{- range .MetricTags }}
		"{{ printf "%s" .}}",
	{{- end }}
	}
	opts := prometheus.GaugeOpts{Name: "{{ .MetricName }}", Help: "{{ index .Provider "help" }}"}
	guage := prometheus.NewGaugeVec(opts, labels).WithLabelValues(
	{{- range .MetricTags }}
		m.{{ printf "%s" .}},
	{{- end }}
	)
	prometheus.MustRegister(guage)
	return {{ .MetricName }}Gauge{
		{{ .MetricName }}: m,
		g: guage,
	}
}

// Set is a wrapper around the prometheus Set(float64) method
func (g {{ .MetricName }}Gauge) Set(s float64) {
	g.g.Set(s)
}
{{end}}
{{- end }}
`

var templates = map[string]*template.Template{
	Prometheus: template.Must(template.New("").Parse(prometheusTemplate)),
}

type MetricFuncs struct {
	MetricName string
	MetricTags []string
	Provider   map[string]string
	Type       string
}

func main() {
	var template *template.Template
	var ok bool

	var p = flag.String("p", "", "metrics provider")
	var output = flag.String("o", "", "output file for generated functions")
	flag.Parse()
	provider := *p

	if template, ok = templates[provider]; !ok {
		log.Fatal("not a valid metrics template")
	}

	funcs := []MetricFuncs{}

	m := reflect.TypeOf(metrics.MetricTypes{})
	for i := 0; i < m.NumField(); i++ {
		varType := m.Field(i).Type
		metricName := varType.Name()

		labels := make([]string, varType.NumField())

		for j := 0; j < varType.NumField(); j++ {
			labels[j] = varType.Field(j).Name
		}

		fun := MetricFuncs{
			MetricName: metricName,
			MetricTags: labels,
			Provider:   map[string]string{},
			Type:       m.Field(i).Tag.Get(TypeStructTag),
		}

		providerTag := m.Field(i).Tag.Get(provider)
		tagattrs := strings.Split(providerTag, ",")
		for _, tagattr := range tagattrs {
			t := strings.Split(tagattr, "=")
			if len(t) == 2 {
				fun.Provider[strings.TrimSpace(t[0])] = strings.TrimSpace(t[1])
			}
		}

		funcs = append(funcs, fun)
	}

	out, err := os.Create(*output)
	if err != nil {
		log.Fatal("Cannot create funcs.go file")
	}
	defer out.Close()

	template.Execute(out, struct {
		Timestamp time.Time
		Funcs     []MetricFuncs
	}{
		Timestamp: time.Now(),
		Funcs:     funcs,
	})
}
