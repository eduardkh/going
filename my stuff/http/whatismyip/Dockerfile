# docker build -t whatismyip .
# docker run --rm -it -p 8080:8080 whatismyip:latest
FROM golang:latest AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/whatismyip
COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
COPY --from=builder /go/bin/whatismyip .
ENTRYPOINT ["./whatismyip"]