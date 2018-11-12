# build stage
FROM golang:1.11.2-alpine3.8 AS builder
RUN apk add --no-cache git openssh gcc musl-dev
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o gotversion cmd/main.go

FROM alpine:latest 
MAINTAINER Andy Lo-A-Foe <andy.loafoe@aemain.com>
WORKDIR /app
COPY --from=builder /build/gotversion /app
RUN apk --no-cache add ca-certificates

VOLUME ["/repo"]
CMD ["/app/gotversion", "/repo"]
