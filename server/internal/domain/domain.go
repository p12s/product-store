package domain

import "time"

// Product
type Product struct {
	Name             string    `csv:"name" json:"name" bson:"name"`
	Price            int       `csv:"price" json:"price" bson:"price"`
	PriceChangeCount int       `csv:"-" json:"price_change_count" bson:"price_change_count"`
	UpdatedAt        time.Time `csv:"-" json:"updated_at" bson:"updated_at"`
	IsUpdated        bool      `csv:"-" json:"is_updated" bson:"is_updated"`
}
