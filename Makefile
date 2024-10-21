build: clean
	@go build -o ./bin/crawler ./main.go

run: build
	@./bin/crawler

clean:
	@rm -rf bin/

test:
	@go test -v ./...
