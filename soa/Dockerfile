FROM golang:1.24.1-bookworm AS builder
FROM alpine:3.18 AS final

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY pkg/ pkg/
COPY cmd/entrypoints/ServicioMenu    cmd/entrypoints/ServicioMenu
COPY cmd/entrypoints/ServicioReserva cmd/entrypoints/ServicioReserva
COPY cmd/entrypoints/ServicioUsuario cmd/entrypoints/ServicioUsuario

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -ldflags="-s -w" -o /bin/ServicioMenu    ./cmd/entrypoints/ServicioMenu && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -ldflags="-s -w" -o /bin/ServicioReserva ./cmd/entrypoints/ServicioReserva && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -ldflags="-s -w" -o /bin/ServicioUsuario ./cmd/entrypoints/ServicioUsuario



# Instalar shell y certificados (si usas HTTPS)
RUN apk add --no-cache ca-certificates bash
WORKDIR /app
COPY --from=builder /bin/ServicioMenu    /bin/ServicioMenu
COPY --from=builder /bin/ServicioReserva /bin/ServicioReserva
COPY --from=builder /bin/ServicioUsuario /bin/ServicioUsuario
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

# Exponer todos los puertos
EXPOSE 8081 8082 8083

ENTRYPOINT ["/entrypoint.sh"]
