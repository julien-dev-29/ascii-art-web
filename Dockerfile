FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ascii-art-web .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/banners ./banners

EXPOSE 8000

CMD ["./ascii-art-web"]
