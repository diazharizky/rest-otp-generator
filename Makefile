# NOTES:
# - Before you use `make run`, make sure you hit `source .env` first to import required environments

.PHONY: generate build clean clean-packr

run:
	go run main.go

migrate:
	go run main.go migrate

test:
	go run main.go test

generate:
	go generate -v ./... && go get -v ./...

build:
	go build -v -o "bin/codebase-rest-golang"

clean:
	rm -rf bin \
	rm -rf api/swagger-spec/swagger.json \
	rm -rf pkg/mariadb/bindata.go \

clean-packr:
	go run github.com/gobuffalo/packr/v2/packr2 clean