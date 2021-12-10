package service

import (
	"context"
	"fmt"

	"github.com/p12s/product-store/server/internal/repository"
	httpTransport "github.com/p12s/product-store/server/internal/transport/http"
	"github.com/p12s/product-store/server/pb"
	"github.com/p12s/product-store/server/pkg/file"
	"github.com/sirupsen/logrus"
)

// Producter
type Producter interface {
	LoadProducts(ctx context.Context, fileSaveDir, url string) error
	GetProducts(ctx context.Context, req *pb.GetProductsRequest) ([]*pb.Product, error)
}

// ProductService
type ProductService struct {
	repo repository.Producter
}

// NewProductService
func NewProductService(repo repository.Producter) *ProductService {
	return &ProductService{repo: repo}
}

// LoadProducts
func (s *ProductService) LoadProducts(ctx context.Context, fileSaveDir, url string) error {
	originalFileName, err := httpTransport.TryDownloadFile(fileSaveDir, url)
	if err != nil {
		logrus.Errorf("download file by url fail: %s", err.Error())
		return fmt.Errorf("download file by url fail: %w", err)
	}
	filePath := fmt.Sprintf("%s/%s", fileSaveDir, originalFileName)

	products, err := file.ReadProducts(filePath)
	if err != nil {
		logrus.Errorf("read products from file fail: %s", err.Error())
		return fmt.Errorf("read products from file fail: %w", err)
	}

	err = s.repo.SaveOrUpdate(ctx, products)
	if err != nil {
		logrus.Errorf("save products to db fail: %s", err.Error())
		return fmt.Errorf("save products to db fail: %w", err)
	}

	err = file.Remove(filePath)
	if err != nil {
		logrus.Errorf("remove file fail: %s", err.Error())
	}

	return nil
}

// GetProducts
func (s *ProductService) GetProducts(ctx context.Context, req *pb.GetProductsRequest) ([]*pb.Product, error) {
	return s.repo.GetProducts(ctx, req)
}
