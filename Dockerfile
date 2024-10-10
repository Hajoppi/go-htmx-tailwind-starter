FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN apk add --update npm
RUN go build -o main ./cmd/main.go
RUN npx tailwindcss \
    -i ./public/css/index.css \
    -o ./public/css/generated.css

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY public ./public
COPY views ./views

CMD ["./main"]
