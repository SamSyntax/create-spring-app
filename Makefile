OUT_DIR := "bin/csa"
ENTRY_POINT := "."
build: main.go
	@go build -o $(OUT_DIR) $(ENTRY_POINT)
run:
	@make build && $(OUT_DIR) 
test:
	@go test -v ./...

