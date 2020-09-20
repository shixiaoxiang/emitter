package web

import (
	"regexp"

	"github.com/emitter-io/emitter/internal/security"
	"net/http"
)

var (
	shortcut = regexp.MustCompile("^[a-zA-Z0-9]{1,2}$")
)

type keygen interface {
	DecryptKey(string) (security.Key, error)
}

// Service represents a self-introspection service.
type Service struct{}

// New creates a new service.
func New() *Service {
	return new(Service)
}

// Login handles a request to create a link.
func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello emitter"))
}
