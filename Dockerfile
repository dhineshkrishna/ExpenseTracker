FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o expense-api ./cmd/server


FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/expense-api /expense-api

USER nonroot:nonroot

EXPOSE 8080

CMD ["/expense-api"]
