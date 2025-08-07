package apiserver

import "fmt"

type APIServer struct {
	config *Config
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (s *APIServer) Start() error {
	fmt.Println("Server is running")
	return nil
}