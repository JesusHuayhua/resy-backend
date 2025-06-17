package main

import (
	"net"
	"net/http"
	"os"
	"soa/services/users/pkg/api/middleware/endpoints"
	"soa/services/users/pkg/api/middleware/transport"
	"soa/services/users/pkg/core/usecase/interfaces"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/grpc/_grpc_test/pb"
	"github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
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

func main() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
		grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	)
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	var (
		service     = interfaces.NuevoServicio()
		eps         = endpoints.NewEndpoints(service)
		httpHandler = transport.NewHTTPHandler(eps)
		grpcServer  = transport.NewGRPCServer(eps)
	)
	var g group.Group
	{
		httpList, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpList, httpHandler)
		}, func(error) {
			httpList.Close()
		})
	}
	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterTestServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}

}
