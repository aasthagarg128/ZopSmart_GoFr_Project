package library

import (
	"github.com/aasthagarg128/ZopSmart_GoFr_Project/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"go.mongodb.org/mongo-driver/bson"
)

type store struct{}

// New is a factory function for the store layer
//
//nolint:revive // library should not be used without proper initialization with required dependency
func New() store {
	return store{}
}

// Get retrieves the details of a book from the library based on the book ID
func (s store) Get(ctx *gofr.Context, id string) (models.Book, error) {
	var book models.Book

	// Fetch the MongoDB collection
	collection := ctx.MongoDB.Collection("books")

	// Create a filter based on the book ID
	filter := bson.M{"id": id}

	// Find the book in the collection
	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return book, errors.DB{Err: err}
	}

	return book, nil
}

// Create adds a new book to the library
func (s store) Create(ctx *gofr.Context, book models.Book) error {
	// Fetch the MongoDB collection
	collection := ctx.MongoDB.Collection("books")

	// Insert the book into the collection
	_, err := collection.InsertOne(ctx, book)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

// Delete removes a book from the library based on the book ID
func (s store) Delete(ctx *gofr.Context, id string) (int, error) {
	// Fetch the MongoDB collection
	collection := ctx.MongoDB.Collection("books")

	// Create a filter based on the book ID
	filter := bson.M{"id": id}

	// Delete the book from the collection
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	return int(result.DeletedCount), nil
}
