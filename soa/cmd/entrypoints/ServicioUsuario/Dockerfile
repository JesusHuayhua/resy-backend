FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copy module files and download deps
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# Copy entrypoint and pkg code
COPY ../../pkg ./pkg
COPY . ./

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /svc

# Stage 2: tiny runtime image
FROM scratch
COPY --from=builder /svc /svc

# If your service listens on 8082, expose it
EXPOSE 8082

ENTRYPOINT ["/svc"]