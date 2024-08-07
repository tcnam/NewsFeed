build:
	go build ./cmd/web/main.go -o ./bin/ecomerce

run: build
	./bin/ecomerce

test:
	go test -v ./...