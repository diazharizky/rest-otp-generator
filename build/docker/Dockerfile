#################
# BUILDER IMAGE #
#################
FROM golang:1.17-alpine AS builder

RUN apk add --no-cache make

WORKDIR /usr/src/reog
COPY . .
RUN make clean clean-packr
RUN make generate && make build

###############
# FINAL IMAGE #
###############
FROM alpine:3.8

RUN apk add --no-cache ca-certificates

# Set config dir
COPY ./configs /etc/reog

COPY --from=builder /usr/src/reog/bin/reog /usr/local/bin/reog

ENTRYPOINT ["/usr/local/bin/reog"]
