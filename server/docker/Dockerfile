FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]