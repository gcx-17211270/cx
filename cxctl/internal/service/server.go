package service

import (
	"net/http"
)

func NewService(c Create, client *http.Client) *Server {
	return &Server{Create: c, client: client}
}

type Server struct {
	Create
	client *http.Client
}
