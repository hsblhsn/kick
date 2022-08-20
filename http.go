package kick

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type HTTPResponseEncoder interface {
	Encode(w *http.ResponseWriter) error
}

type HTTPRequestDecoder interface {
	Decode(r *http.Request) error
}

func UnmarshalJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func MarshalJSON(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = MarshalJSON(w, map[string]string{
		"error": err.Error(),
	})
}

type Server struct {
	server *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
