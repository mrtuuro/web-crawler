BASE_URL = "https://wagslane.dev"
	MAX_CONCURRENCY = 3
	MAX_PAGES = 30


build: clean
	@go build -o ./bin/crawler ./*.go

run: build
	@./bin/crawler $(BASE_URL) $(MAX_CONCURRENCY) $(MAX_PAGES)

clean:
	@go clean -cache
	@go clean -testcache
	@rm -rf bin/

test: clean
	@go test -v ./...
