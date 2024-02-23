package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	Timeout = 60 * time.Second
)

func ConnectGRPC(host string, port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
