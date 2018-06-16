clean:
	rm -f metrics/funcs.go
run: clean
	cd metrics && go generate
	go build .
	./genericmetrics
