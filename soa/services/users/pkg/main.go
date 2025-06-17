package main

import (
	"net"
	"os"
	"soa/services/users/pkg/api/middleware/endpoints"
	"soa/services/users/pkg/api/middleware/transport"

	"github.com/go-kit/log"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func main_test() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
		grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	)
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	var (
		service     = interfaces.nuevoServicio()
		eps         = endpoints.NewEndpoints(service)
		httpServer  = transport.NewHTTPHandler(eps)
		httpHandler = transport.NewGRPCServer(eps)
	)

}
