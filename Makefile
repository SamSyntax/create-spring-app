build: internal/cmd/main.go
	@go build -o bin/csa internal/cmd/main.go
test:
	@go test -v ./...

