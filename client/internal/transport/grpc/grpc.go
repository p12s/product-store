package grpc

import (
	"fmt"

	"github.com/p12s/product-store/client/internal/config"
	"github.com/p12s/product-store/client/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client
type Client struct {
	Service pb.ProductServiceClient
	conn    *grpc.ClientConn
}

// New
func New(cfg config.Config) (*Client, error) {
	conn, err := grpc.Dial(cfg.Server.Host, grpc.WithInsecure()) // nolint
	if err != nil {
		logrus.Fatalf("error grpc client: %s\n", err.Error())
		return nil, fmt.Errorf("error grpc client: %w\n", err)
	}

	return &Client{
		conn:    conn,
		Service: pb.NewProductServiceClient(conn),
	}, nil
}

// Close
func (c *Client) Close() {
	c.conn.Close()
}
