default: dist

dist: clean
	mkdir -p pkg
	mkdir -p pkg/linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pkg/linux-amd64/log-me

	mkdir -p pkg/darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o pkg/darwin-amd64/log-me

clean:
	rm -rf pkg
	rm -f log-me
