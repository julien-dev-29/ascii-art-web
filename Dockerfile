FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/ascii-art-web .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/banners ./banners

LABEL maintainer="ton_nom"
LABEL version="1.0"
LABEL description="Ascii Art Web Server in Go"

RUN adduser -D appuser
USER appuser

EXPOSE 8000

CMD ["./ascii-art-web"]
