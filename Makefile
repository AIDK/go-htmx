build:
	@echo "building go-htmx"
	@go build -o bin/go-htmx main.go
	@echo "build done"

run: build
	@echo "running go-htmx"
	@./bin/go-htmx
