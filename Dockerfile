ARG VERSION=1.0.0

FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ascii-art-web .

FROM alpine:3.19

ARG VERSION
ARG BUILD_DATE

LABEL org.opencontainers.image.title="ascii-art-web"
LABEL org.opencontainers.image.description="Web-based ASCII art generator with multiple banner styles"
LABEL org.opencontainers.image.version="${VERSION}"
LABEL org.opencontainers.image.authors="Julien"
LABEL org.opencontainers.image.created="${BUILD_DATE}"
LABEL org.opencontainers.image.url="https://github.com/01edu/go/ascii-art-web"
LABEL org.opencontainers.image.source="https://github.com/01edu/go/ascii-art-web"
LABEL org.opencontainers.image.licenses="MIT"

RUN adduser -D -h /app appuser

WORKDIR /app

COPY --from=builder --chown=appuser:appuser /app/ascii-art-web .
COPY --from=builder --chown=appuser:appuser /app/templates ./templates
COPY --from=builder --chown=appuser:appuser /app/banners ./banners
COPY --from=builder --chown=appuser:appuser /app/static ./static

USER appuser

EXPOSE 8000

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8000/ || exit 1

CMD ["./ascii-art-web"]
