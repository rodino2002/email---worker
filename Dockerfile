# ---------- BUILD ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Copia tudo da tua pasta local (incluindo o .env) para /app no contentor
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o email-worker ./cmd/main.go

# ---------- RUN ----------
FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache tzdata
ENV TZ=Africa/Luanda

# Copia o binário do builder para /app/
COPY --from=builder /app/email-worker .

# Copia o .env do builder para /app/
COPY --from=builder /app/.env .

RUN mkdir -p /app/logs && \
    adduser -D appuser && \
    chown -R appuser:appuser /app

USER appuser

CMD ["./email-worker"]