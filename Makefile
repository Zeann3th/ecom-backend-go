build:
	@echo "Building Application..."
	@go build -o bin/app cmd/main.go
test:
	@echo "Running Tests..."
	@go test -v ./...
clean:
	@echo "Cleaning Up..."
	@rm -rf bin
run: build
	@echo "Running Application..."
	@./bin/app

