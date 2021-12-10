package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/p12s/product-store/server/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server
type Server struct {
	grpcSrv        *grpc.Server
	productService pb.ProductServiceServer
}

// New
func New(productService pb.ProductServiceServer) *Server {
	return &Server{
		grpcSrv:        grpc.NewServer(),
		productService: productService,
	}
}

// ListenAndServe
func (s *Server) ListenAndServe(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatalf("error grpc listening: %s\n", err.Error())
	}

	pb.RegisterProductServiceServer(s.grpcSrv, s.productService)
	reflection.Register(s.grpcSrv)

	go func() {
		if err := s.grpcSrv.Serve(listen); err != nil {
			logrus.Fatalf("error while running grpc server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("Server shutting down. %s", time.Now().Format(time.RFC3339))
	s.grpcSrv.GracefulStop()
}
