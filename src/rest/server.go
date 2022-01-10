package rest

import (
	"encoding/json"
	"fmt"
	"kmsbot/domain"
	"net/http"
)

type Server struct {
	*http.Server
	core *domain.Core
}

func NewServer(port string, core *domain.Core) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: nil,
		},
		core: core}
}

func (s *Server) SetRoutes() {
	r := http.NewServeMux()
	r.HandleFunc("/send", s.CorsHandler(s.SendHandler()))

	s.Handler = r
}

func (s *Server) SendHandler() http.Handler {
	type ReqData struct {
		Text string `json:"text"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ReqData

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)

			return
		}

		if req.Text != "" {
			s.core.SendNotification(req.Text)
		}
	})
}

func (s *Server) CorsHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	}
}
