package server

import (
	"fmt"
	"net"
)

type GRPCServer struct {
	listener net.Listener
}

func New(listenPort int) *GRPCServer {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		fmt.Println("ошибка создания входного сокета ", err)
		return nil
	}
	return &GRPCServer{listener: l}
}

func (s *GRPCServer) Run() {
}
