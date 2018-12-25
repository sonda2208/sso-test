package app

import (
	"github.com/sonda2208/sso-test/model"
)

type Server struct {
	config *model.ServiceSettings
}

func New(conf *model.ServiceSettings) (*Server, error) {
	s := &Server{
		config: conf,
	}

	return s, nil
}

func (s *Server) Config() model.ServiceSettings {
	return *s.config
}
