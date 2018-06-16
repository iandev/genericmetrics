clean:
	rm -f metrics/funcs.go
build: clean
	go generate ./...
	go build .
run: clean build
	./genericmetrics
