build: clean
	@go build -o ./bin/crawler ./main.go

run: build
	@./bin/crawler

clean:
	@go clean -cache
	@go clean -testcache
	@rm -rf bin/

test: clean
	@go test -v ./...
