package handlers

import (
	"fmt"

	"github.com/aasthagarg128/ZopSmart_GoFr_Project/models"
	"github.com/aasthagarg128/ZopSmart_GoFr_Project/stores"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store stores.Library
}

// New is a factory function for the handler layer
//
//nolint:revive // handler should not be used without proper initialization with required dependency
func New(l stores.Library) handler {
	return handler{store: l}
}

func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
	// Retrieve the book ID from the request parameters
	bookID := ctx.Param("id")

	// Call the store's Get method to fetch the book details
	book, err := h.store.Get(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var book models.Book

	// Bind the request body to the book model
	err := ctx.Bind(&book)
	if err != nil {
		return nil, err
	}

	// Call the store's Create method to add the book to the library
	err = h.store.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	return "New book added!", nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	// Retrieve the book ID from the request parameters
	bookID := ctx.Param("id")

	// Call the store's Delete method to remove the book from the library
	deleteCount, err := h.store.Delete(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("%v books deleted!", deleteCount), nil
}

func (h handler) Ping(ctx *gofr.Context) (interface{}, error) {
	return "pong", nil
}