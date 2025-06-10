package ports

import (
	"context"
	"net/http"
	"os"
	"soa/services/users/cmd/core"

	log "github.com/go-kit/log"
)

type Servicio1 struct{}

var logger log.Logger

func nuevoServicio() core.Service {
	return &Servicio1{}
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

func (s1 *Servicio1) Get(_ context.Context, userId string) {

}

func (s2 *Servicio1) Status(_ context.Context, userId string) {

}

func (s3 *Servicio1) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking status")
	return http.StatusOK, nil
}
