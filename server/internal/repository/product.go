package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/p12s/product-store/server/internal/domain"
	"github.com/p12s/product-store/server/pb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DEFAULT_PRICE_CHANGE_COUNT = 1
	IS_UPDATED_FLAG            = true
)

// Producter
type Producter interface {
	SaveOrUpdate(ctx context.Context, products []domain.Product) error
	GetProducts(ctx context.Context, req *pb.GetProductsRequest) ([]*pb.Product, error)
}

// Product
type Product struct {
	products *mongo.Collection
}

// NewProduct - constructor
func NewProduct(products *mongo.Collection) *Product {
	return &Product{products: products}
}

// SaveOrUpdate
func (r *Product) SaveOrUpdate(ctx context.Context, products []domain.Product) error {
	// every product
	for _, newProduct := range products {
		newProduct.UpdatedAt = time.Now()
		newProduct.PriceChangeCount = DEFAULT_PRICE_CHANGE_COUNT
		newProduct.IsUpdated = IS_UPDATED_FLAG

		// create if not exists
		existsProduct, findErr := r.FindOneByName(ctx, newProduct.Name)
		if errors.Is(findErr, domain.ErrProductNotFound) {
			_, insertErr := r.products.InsertOne(ctx, newProduct)
			if insertErr != nil {
				return fmt.Errorf("save product fail: %w/n", insertErr)
			}
			continue
		}
		if findErr != nil {
			return fmt.Errorf("find product by name other fail: %w/n", findErr)
		}

		// updted fileds depends of price changes
		var updateBson primitive.M
		if newProduct.Price != existsProduct.Price {
			updateBson = bson.M{"$set": bson.M{
				"price":              newProduct.Price,
				"price_change_count": existsProduct.PriceChangeCount + DEFAULT_PRICE_CHANGE_COUNT,
				"updated_at":         newProduct.UpdatedAt,
				"is_updated":         IS_UPDATED_FLAG,
			}}
		} else {
			updateBson = bson.M{"$set": bson.M{
				"updated_at": newProduct.UpdatedAt,
				"is_updated": IS_UPDATED_FLAG,
			}}
		}

		// update
		_, updateErr := r.products.UpdateOne(ctx, bson.M{"name": newProduct.Name}, updateBson)
		if updateErr != nil {
			return fmt.Errorf("update product by name fail: %w/n", updateErr)
		}
	}

	// delete no longer existing in the store
	res, deleteErr := r.products.DeleteMany(ctx, bson.M{"is_updated": bson.M{"$exists": false}})
	if deleteErr != nil {
		return fmt.Errorf("delete not updated fail: %w/n", deleteErr)
	} else {
		logrus.Println("non-existent products removing count:", res.DeletedCount)
	}

	_, updateErr := r.products.UpdateMany(ctx, bson.M{"is_updated": IS_UPDATED_FLAG}, bson.M{"$unset": bson.M{"is_updated": IS_UPDATED_FLAG}})
	if updateErr != nil {
		return fmt.Errorf("update flag remove fail: %w/n", updateErr)
	}

	return nil
}

// GetProducts
func (r *Product) GetProducts(ctx context.Context, req *pb.GetProductsRequest) ([]*pb.Product, error) {
	opts := getPaginationOpts(&domain.PaginationQuery{
		Skip:  req.GetSkip(),
		Limit: req.GetLimit(),
	})
	opts.SetSort(bson.M{"price": toInt(req.GetPriceOrder())})

	cur, err := r.products.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("product find fail: %w/n", err)
	}

	var products []domain.Product
	if err := cur.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("product convert fail: %w/n", err)
	}

	return domainToProto(products), err
}

func (r *Product) DeleteMany(ctx context.Context, filter primitive.M) error {
	_, err := r.products.DeleteMany(ctx, filter)
	return err
}

// FindOneByName
func (r *Product) FindOneByName(ctx context.Context, name string) (domain.Product, error) {
	var product domain.Product
	err := r.products.FindOne(ctx, bson.M{"name": name}).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Product{}, domain.ErrProductNotFound
		}
		return domain.Product{}, fmt.Errorf("product find fail: %w/n", err)
	}
	return product, err
}

// toInt
func toInt(sort pb.SortOrder) int {
	var sortNumber int
	switch sort {
	case pb.SortOrder_ASC:
		sortNumber = 1
	case pb.SortOrder_DESC:
		sortNumber = -1
	}
	return sortNumber
}

// domainToProto
func domainToProto(products []domain.Product) []*pb.Product {
	result := make([]*pb.Product, len(products))
	for i := 0; i < len(products); i++ {
		result[i] = &pb.Product{
			Name:             products[i].Name,
			Price:            int64(products[i].Price),
			PriceChangeCount: int64(products[i].PriceChangeCount),
			UpdatedAt:        timestamppb.New(products[i].UpdatedAt),
		}
	}
	return result
}
