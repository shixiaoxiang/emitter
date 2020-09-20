package broker

import (
	"net/http"
)

// AddWebHandle  is ok
func AddWebHandle(s *Service, mux *http.ServeMux) {
	mux.HandleFunc("/emitter/login/", s.web.Login)
}
