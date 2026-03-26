build:
	go build -o drs main.go

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
