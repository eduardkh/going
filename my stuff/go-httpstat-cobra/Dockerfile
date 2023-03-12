# docker build -t go-httpstat-cobra .
# docker run --rm go-httpstat-cobra:latest
FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/go-httpstat-cobra
COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
# for TLS connections to work
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/go-httpstat-cobra .
ENTRYPOINT ["./go-httpstat-cobra"]
