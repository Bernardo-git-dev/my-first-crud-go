# Fase 1 - Build
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && \
  go build -o main . && \
  chmod +x main

# Fase 2 - Execução
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 3333

CMD ["./main"]
