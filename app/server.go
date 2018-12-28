package app

import (
	"github.com/sonda2208/sso-test/model"
	"github.com/sonda2208/sso-test/store"
	"github.com/sonda2208/sso-test/store/mem_store"
)

type Server struct {
	config *model.ServiceSettings
	store  store.TokenStore
}

func New(conf *model.ServiceSettings) (*Server, error) {
	s := &Server{
		config: conf,
		store:  mem_store.NewOnMemoryTokenStore(),
	}

	return s, nil
}

func (s *Server) Config() model.ServiceSettings {
	return *s.config
}
