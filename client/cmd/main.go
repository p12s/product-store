package main

import (
	"context"
	"fmt"
	"io"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/p12s/product-store/client/internal/config"
	"github.com/p12s/product-store/client/internal/transport/grpc"
	"github.com/p12s/product-store/client/pb"
	"github.com/sirupsen/logrus"
)

const (
	PAGE_PRODUCT_LIMIT        = 5
	FAKE_PAGE_NUMBER          = 10
	FAKE_PAGE_SCROLL_DURATION = 2 * time.Second
)

// @title Store app gRPC-client
// @version 0.0.1
// @description Simple application for loading/getting products
func main() {
	runtime.GOMAXPROCS(1)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error reading env variables from file: %s\n", err.Error())
	}
	cfg, err := config.New()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err.Error())
	}

	storeClient, err := grpc.New(*cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize grpc client: %s\n", err.Error())
	}
	defer storeClient.Close()

	// 1. load product by url
	response, err := storeClient.Service.LoadProducts(context.Background(),
		&pb.LoadProductsRequest{Url: cfg.Store.Url})
	if err != nil {
		logrus.Fatalf("failed to load products:\n %s\n%s\n", cfg.Store.Url, err.Error())
	}
	fmt.Printf("download product response:\n%s\n\n", response.GetMessage())

	// 2. getting by page
	for page := 1; page <= 2; page++ {
		gettingProductsByPage(storeClient, int64(page))
	}

	// 3. "infinite" getting
	infiniteGettingProducts(storeClient)
}

// gettingProductsByPage
func gettingProductsByPage(storeClient *grpc.Client, pageNumber int64) {
	var skip int64 = 0
	if pageNumber > 1 {
		skip = PAGE_PRODUCT_LIMIT*(pageNumber-1) - 1 // 0 4 9 14
	}

	products, err := storeClient.Service.GetProducts(
		context.Background(),
		&pb.GetProductsRequest{
			Limit:      PAGE_PRODUCT_LIMIT,
			Skip:       skip,
			PriceOrder: pb.SortOrder_ASC, // sort prices in ascending order
		},
	)
	if err != nil {
		logrus.Errorf("failed to load %d page products: %s\n", pageNumber, err.Error())
	}

	printProducts(pageNumber, products.Product)
}

// infiniteGettingProducts
func infiniteGettingProducts(storeClient *grpc.Client) {
	stream, err := storeClient.Service.GetProductsInfinite(context.Background())
	if err != nil {
		logrus.Errorf("getting products fail: %s\n", err.Error())
	}

	waitCh := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitCh)
				return
			}
			if err != nil {
				logrus.Fatalf("receive note fail: %v", err)
			}

			printProducts(FAKE_PAGE_NUMBER, in.GetProduct())
		}
	}()

	for _, id := range []int{1, 2, 3} {
		if err := stream.Send(&pb.GetProductsRequest{
			Limit:      PAGE_PRODUCT_LIMIT,
			Skip:       int64(id),
			PriceOrder: pb.SortOrder_ASC,
		}); err != nil {
			logrus.Fatalf("send note fail: %v", err)
		}
		time.Sleep(FAKE_PAGE_SCROLL_DURATION)
	}
	err = stream.CloseSend()
	if err != nil {
		logrus.Fatalf("close stream fail: %v", err)
	}
	<-waitCh
}

// printProducts
func printProducts(pageNumber int64, list []*pb.Product) {
	fmt.Printf("\n%d page products:\n", pageNumber)
	for _, item := range list {
		fmt.Println(item)
	}
}
