FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o server .

FROM alpine:latest as runner

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]