# ---------- Build Stage ----------
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Fully static binary (no CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o expense-api ./cmd/server

# ---------- Runtime Stage ----------
FROM gcr.io/distroless/static:nonroot

WORKDIR /

# Copy binary
COPY --from=builder /app/expense-api /expense-api

# Run as non-root (secure + lightweight)
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/expense-api"]