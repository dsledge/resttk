package resttk

import (
	"bytes"
	"errors"
	"net/http"
)

type Server struct {
	Addr   string
	Port   string
	Cert   string
	Key    string
	SSL    bool
	Routes *Router
}

func (s *Server) Run() error {
	http.Handle("/", s.Routes)

	if s.SSL {
		if err := http.ListenAndServeTLS(s.buildURL(), s.Cert, s.Key, nil); err != nil {
			return err
		}
		return nil
	} else {
		if err := http.ListenAndServe(s.buildURL(), nil); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Secure parameter must be specified for the resttk instance")
}

func (s *Server) buildURL() string {
	var buffer bytes.Buffer
	if len(s.Addr) > 0 {
		buffer.WriteString(s.Addr)
	}
	buffer.WriteString(":")
	if len(s.Port) > 0 {
		buffer.WriteString(s.Port)
	} else {
		buffer.WriteString("9001")
	}
	return buffer.String()
}
