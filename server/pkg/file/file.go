package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jszwec/csvutil"
	"github.com/p12s/product-store/server/internal/domain"
)

const READER_COMMA = ';'

// ReadProducts - reading products array from csv-file
func ReadProducts(filePath string) ([]domain.Product, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file fail: %w/n", err)
	}
	reader := csv.NewReader(csvFile)
	reader.Comma = READER_COMMA

	fileExt := filepath.Ext(filePath)
	if len(fileExt) > 1 {
		fileExt = fileExt[1:]
	}

	userHeader, _ := csvutil.Header(domain.Product{}, fileExt)
	dec, err := csvutil.NewDecoder(reader, userHeader...)
	if err != nil {
		return nil, fmt.Errorf("file header decode fail: %w/n", err)
	}

	var products []domain.Product
	for {
		var p domain.Product
		if err := dec.Decode(&p); err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("product decode fail: %w/n", err)
		}
		products = append(products, p)
	}

	return products, nil
}

// Remove - remove file
func Remove(filePath string) error {
	return os.Remove(filePath)
}
