BASE_URL = "https://wagslane.dev"


build: clean
	@go build -o ./bin/crawler ./*.go

run: build
	@./bin/crawler $(BASE_URL)

clean:
	@go clean -cache
	@go clean -testcache
	@rm -rf bin/

test: clean
	@go test -v ./...
