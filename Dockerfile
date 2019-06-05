FROM golang:1.12 as builder

WORKDIR /go/src/github.com/srvaroa/jsonrouter
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -o jsonrouter ./cmd/main.go

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/srvaroa/jsonrouter/jsonrouter /jsonrouter
COPY --from=builder /go/src/github.com/srvaroa/jsonrouter/samples/config.json /config.json
CMD ["/jsonrouter", "-routes", "/config.json"]
