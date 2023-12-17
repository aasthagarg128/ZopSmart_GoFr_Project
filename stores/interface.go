package stores

import (
	"ZopSmart_GoFr_Project/models"
	"gofr.dev/pkg/gofr"
)

// Library defines the methods for interacting with the library
type Library interface {
	Get(ctx *gofr.Context, id string) (models.Book, error)
	Create(ctx *gofr.Context, book models.Book) error
	Delete(ctx *gofr.Context, id string) (int, error)
}
