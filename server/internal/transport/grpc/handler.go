package grpc

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/p12s/product-store/server/internal/service"
	"github.com/p12s/product-store/server/pb"
	"github.com/sirupsen/logrus"
)

// Handler
type Handler struct {
	services    *service.Service
	fileSaveDir string
}

// NewHandler - constructor
func NewHandler(services *service.Service, fileSaveDir string) *Handler {
	return &Handler{services: services, fileSaveDir: fileSaveDir}
}

// LoadProducts - getting csv-file by from external source(url) and save products
func (h *Handler) LoadProducts(ctx context.Context, req *pb.LoadProductsRequest) (*pb.LoadProductsResponse, error) {
	err := h.services.LoadProducts(ctx, h.fileSaveDir, req.GetUrl())
	if err != nil {
		logrus.Errorf("load products fail: %s", err.Error())
		return nil, fmt.Errorf("load products fail: %w", err)
	}

	return &pb.LoadProductsResponse{
		Code:    http.StatusOK,
		Message: "Products downloaded OK",
	}, nil
}

// GetProducts - getting products from local db
func (h *Handler) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	products, err := h.services.GetProducts(ctx, req)
	if err != nil {
		logrus.Errorf("get products fail: %s", err.Error())
		return nil, fmt.Errorf("get products fail: %w", err)
	}

	return &pb.GetProductsResponse{Product: products}, nil
}

// GetProductsInfinite - "infinite" getting products from local db
func (h *Handler) GetProductsInfinite(stream pb.ProductService_GetProductsInfiniteServer) error {
	ctx := context.Background()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logrus.Errorf("error while reading client stream %s", err.Error())
		}

		products, err := h.services.GetProducts(ctx, req)
		if err != nil {
			logrus.Errorf("error while infinite getting products %s", err.Error())
		}

		err = stream.Send(&pb.GetProductsResponse{
			Product: products,
		})
		if err != nil {
			logrus.Errorf("error while sending data to client %s", err.Error())
		}
	}
}
