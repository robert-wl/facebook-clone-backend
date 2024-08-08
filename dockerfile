FROM golang:alpine3.20

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o /app/binary

EXPOSE 8080:8080

ENTRYPOINT ["/app/binary"]