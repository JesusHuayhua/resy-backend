FROM golang:1.24.4-bookworm AS builder
WORKDIR /app 
COPY . . 
RUN if [ ! -f go.mod ]; then \
      echo "[GO MOD] go.mod not found, initializing…"; \
      go mod init soa; \
    fi && \
    echo "[GO MOD] running go mod tidy…" && \
    go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -gcflags="all=-N -l" -o /bin/ServicioMenu    ./cmd/entrypoints/ServicioMenu && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -gcflags="all=-N -l" -o /bin/ServicioReserva ./cmd/entrypoints/ServicioReserva && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -gcflags="all=-N -l" -o /bin/ServicioUsuario ./cmd/entrypoints/ServicioUsuario

FROM alpine:3.18 AS final
#RUN apk add --no-cache ca-certificates bash
WORKDIR /app
COPY --from=builder /bin/ServicioMenu    /bin/ServicioMenu
COPY --from=builder /bin/ServicioReserva /bin/ServicioReserva
COPY --from=builder /bin/ServicioUsuario /bin/ServicioUsuario

# Genera entrypoint.sh al momento que entrypoint.sh cuando se invoke a dockerfile.
RUN echo '#!/bin/sh' > /entrypoint.sh && \
    echo 'echo "Iniciando ServicioMenu..."' >> /entrypoint.sh && \
    echo '/bin/ServicioMenu &' >> /entrypoint.sh && \
    echo 'echo "Iniciando ServicioReserva..."' >> /entrypoint.sh && \
    echo '/bin/ServicioReserva &' >> /entrypoint.sh && \
    echo 'echo "Iniciando ServicioUsuario..."' >> /entrypoint.sh && \
    echo '/bin/ServicioUsuario &' >> /entrypoint.sh && \
    echo 'wait' >> /entrypoint.sh && \
    chmod +x /entrypoint.sh
	
EXPOSE 8081 8082 8083
ENTRYPOINT ["/entrypoint.sh"]
