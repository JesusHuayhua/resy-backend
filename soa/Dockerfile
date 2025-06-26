FROM golang:1.24.1-bookworm AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN if [ ! -f go.mod ]; then \
      echo "[INFO] no go.mod, init module..."; \
      go mod init soa; \
    fi \
 && go mod tidy

FROM base AS builder-menu
WORKDIR /app
COPY pkg ./pkg
COPY cmd/entrypoints/ServicioMenu ./cmd/entrypoints/ServicioMenu
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/ServicioMenu ./cmd/entrypoints/ServicioMenu


FROM base AS builder-user
WORKDIR /app
COPY pkg ./pkg
COPY cmd/entrypoints/ServicioUsuario ./cmd/entrypoints/ServicioUsuario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/ServicioUsuario ./cmd/entrypoints/ServicioUsuario


FROM scratch
COPY --from=builder-menu /bin/ServicioMenu /bin/ServicioMenu
COPY --from=builder-user /bin/ServicioUsuario /bin/ServicioUsuario 