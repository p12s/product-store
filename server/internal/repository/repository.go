package repository

import (
	"github.com/p12s/product-store/server/internal/domain"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const DB_NAME = "store"
const PRODUCT_COLLECTION_NAME = "product"

// Repository
type Repository struct {
	Producter
}

// NewRepository
func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Producter: NewProduct(db.Database(DB_NAME).Collection(PRODUCT_COLLECTION_NAME)),
	}
}

// getPaginationOpts
func getPaginationOpts(pagination *domain.PaginationQuery) *options.FindOptions {
	var opts *options.FindOptions
	if pagination != nil {
		opts = &options.FindOptions{
			Skip:  &pagination.Skip,
			Limit: &pagination.Limit,
		}
	}

	return opts
}
