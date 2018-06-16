clean:
	rm -f metrics/funcs.go
run: clean
	go generate
	go build .
	./genericmetrics
