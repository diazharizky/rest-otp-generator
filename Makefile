# NOTES:
# - Before you use `make run`, make sure you hit `source .env` first to import required environments

.PHONY: generate build clean clean-packr

run:
	go run main.go

generate:
	go generate -v ./... && go get -v ./...

test:
	go test ./...

build:
	go build -v -o "bin/rest-otp-generator"

clean:
	rm -rf bin \
	rm -rf api/swagger-spec/swagger.json \

clean-packr:
	go run github.com/gobuffalo/packr/v2/packr2 clean