FROM golang:alpine3.20 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o /app/binary

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/binary /app/binary

EXPOSE 8080

ENTRYPOINT ["/app/binary"]
