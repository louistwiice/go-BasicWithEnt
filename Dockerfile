FROM golang:1.18-alpine

WORKDIR /go/src/code

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o app api/*.go

CMD ["./app"]
