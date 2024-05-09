FROM golang:1.21.1-alpine

RUN go version

ENV GOPATH=/

COPY . .

RUN go mod download

RUN go build -o app ./cmd

EXPOSE 8000

CMD ["./app"]