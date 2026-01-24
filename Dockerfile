# docker build -t jsonplaceholder-clone:1.0.2 .
# docker save jsonplaceholder-clone:1.0.2 -o jsonplaceholder-clone_1.0.2.tar
# docker load -i jsonplaceholder-clone_1.0.2.tar

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server"]
