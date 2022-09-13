build:
	go build

test:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...

run:
	go run .

lint:
	golangci-lint run --out-format checkstyle --issues-exit-code 0 > golangci-report.out

generate:
	go generate ./...
