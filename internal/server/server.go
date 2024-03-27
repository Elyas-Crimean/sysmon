package server

import (
	"fmt"
	"net"
	"time"

	"github.com/Elyas-Crimean/sysmon/api"
	"github.com/Elyas-Crimean/sysmon/internal/storage"
	"google.golang.org/grpc"
)

type Server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	storage    *storage.Storage
}

func New(listenPort int) *Server {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		fmt.Println("ошибка создания входного сокета ", err)
		return nil
	}
	return &Server{listener: l}
}

func (s *Server) Run(storage *storage.Storage) {
	s.storage = storage
	s.grpcServer = grpc.NewServer()
	api.RegisterSysmonServer(s.grpcServer, s)
}

func (s *Server) ProbeQuery(q *api.Query, pqServer api.Sysmon_ProbeQueryServer) error {
	ticker := time.NewTicker(q.Interval.AsDuration())
	stopCH := pqServer.Context().Done()
sessionLoop:
	for {
		data := &api.Data{}
		for k, v := range s.storage.GetAvg(q.Window.AsDuration()) {
			data.Probe = append(data.Probe, &api.Probe{
				Key:   k,
				Value: v,
			})
		}
		pqServer.Send(data)
		select {
		case <-ticker.C:
		case <-stopCH:
			break sessionLoop
		}
	}
	return nil
}
