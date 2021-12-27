# NOTES:
# - Before you use `make run`, make sure you hit `source .env` first to import required environments

.PHONY: generate build clean clean-packr

run:
	go run cmd/main.go

test:
	go clean -testcache && \
	go test ./... -v

generate:
	go generate -v ./... && go get -v ./...

build:
	go build -v -o "bin/rog"

clean:
	rm -rf bin \
	rm -rf api/swagger-spec/swagger.json \

clean-packr:
	go run github.com/gobuffalo/packr/v2/packr2 clean
