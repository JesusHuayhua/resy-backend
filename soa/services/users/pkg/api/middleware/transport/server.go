package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"soa/services/users/pkg/api/middleware/endpoints"
	"soa/services/users/pkg/core/response"
	"soa/services/users/pkg/core/svc_internal"
	"strconv"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

/*
	Documentacion:
	1. service_status, devuelve el estado del servicio.
	2. status
*/

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/service_status", httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))
	m.Handle("/status", httptransport.NewServer(
		ep.StatusEndpoint,
		decodeHTTPStatusRequest,
		encodeResponse,
	))
	m.Handle("/get", httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))
	m.Handle("/users", httptransport.NewServer(
		ep.UsuarioEndpoint,
		decodeHTTPUsersRequest,
		encodeResponse,
	))

	m.Handle("/roles", httptransport.NewServer(
		ep.RolesEndpoint,
		decodeHTTPRolesRequest,
		encodeResponse,
	))

	return m
}

/*

// Nota: Mejorar esto con funciones genericas.
func parseInput1(req *response.RolesRequest, r *http.Request) (interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	q := r.URL.Query()
	opStr := q.Get("tipo_op")
	if opStr == "" {
		return nil, errors.New("falta el parámetro tipo_op")
	}
	op, err := strconv.Atoi(opStr)
	if err != nil {
		return nil, errors.New("tipo_op debe ser entero")
	}
	req.TipoOp = op
	for k, vals := range q {
		if k == "tipo_op" {
			continue
		}
		for _, v := range vals {
			req.Args = append(req.Args, svc_internal.Filter{
				Key:   k,
				Value: v,
			})
		}
	}
	return req, nil
}

func parseInput2(req *response.UsuarioRequest, r *http.Request) (interface{}, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	q := r.URL.Query()
	opStr := q.Get("tipo_op")
	if opStr == "" {
		return nil, errors.New("falta el parámetro tipo_op")
	}
	op, err := strconv.Atoi(opStr)
	if err != nil {
		return nil, errors.New("tipo_op debe ser entero")
	}
	req.TipoOp = op
	for k, vals := range q {
		if k == "tipo_op" {
			continue
		}
		for _, v := range vals {
			req.Args = append(req.Args, svc_internal.Filter{
				Key:   k,
				Value: v,
			})
		}
	}
	return req, nil
}
*/

type HasBaseReq interface {
	~struct {
		TipoOp int
		Args   []svc_internal.Filter
	}
}

func parseInput(req response.GenericRequest, r *http.Request) (response.GenericRequest, error) {
	if err := r.ParseForm(); err != nil {
		return req, err
	}

	q := r.URL.Query()

	opStr := q.Get("tipo_op")
	if opStr == "" {
		return req, errors.New("falta el parámetro tipo_op")
	}
	op, err := strconv.Atoi(opStr)
	if err != nil {
		return req, errors.New("tipo_op debe ser entero")
	}
	req.TipoOp = op // OK: campo visible
	for k, vals := range q {
		if k == "tipo_op" {
			continue
		}
		for _, v := range vals {
			req.Args = append(req.Args,
				svc_internal.Filter{Key: k, Value: v}) // OK
		}
	}
	return req, nil
}

func decodeHTTPRolesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req response.GenericRequest
	req_res, _ := parseInput(req, r)
	return req_res, nil
}

// Dos formas: JSON o URL, esto es lo de menos, para facilidad de testing, usare URL schema.
func decodeHTTPUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req response.GenericRequest
	/*
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return nil, err
		}
	*/
	req_res, _ := parseInput(req, r)
	return req_res, nil
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req response.GetRequest
	if r.ContentLength == 0 {
		logger.Log("Get request with no body")
		return req, nil
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req response.StatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req response.ServiceStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case svc_internal.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case svc_internal.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
