FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .env .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go mod tidy

RUN go build -o /app/binary

EXPOSE 8080:8080

ENTRYPOINT ["/app/binary"]