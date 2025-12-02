package rest

import (
	"ecommerce/config"
	"fmt"
)

type Server struct {
	cnf 	*config.Config
}

func NewServer(cnf *config.Config) *Server {
	return &Server{
		cnf:   cnf,
	}
}

func Start() {
	fmt.Println("Server starting... (temporary mode)")
	// manager := middleware.newManager()
}