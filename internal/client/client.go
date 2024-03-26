package client

import (
	"context"
	"time"

	"github.com/Elyas-Crimean/sysmon/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	sysmon   api.SysmonClient
	interval time.Duration
	window   time.Duration
}

func NewClient(host string, interval time.Duration, window time.Duration) *GRPCClient {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer ctxCancel()
	conn, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil
	}
	grpcClient := &GRPCClient{
		interval: interval,
		window:   window,
	}
	grpcClient.sysmon = api.NewSysmonClient(conn)
	return grpcClient
}
