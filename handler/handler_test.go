package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/ZopSmart_GoFr_Project/models"
	"gofr.dev/ZopSmart_GoFr_Project/stores"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
)

func initializeHandlersTest(t *testing.T) (*stores.MockLibrary, handler, *gofr.Gofr) {
	ctrl := gomock.NewController(t)

	store := stores.NewMockLibrary(ctrl) // Change from MockCustomer to MockLibrary
	h := New(store)
	app := gofr.New()

	return store, h, app
}

func TestHandler_Get(t *testing.T) {
	tests := []struct {
		desc        string
		queryParams string
		bookID      string
		resp        interface{}
		err         error
	}{
		{"get without params", "", "123", models.Book{ID: "123", Title: "Sample Book", Author: "John Doe", Publisher: "Publisher A"}, nil},
		{"get with valid bookID", "id=123", "123", models.Book{ID: "123", Title: "Sample Book", Author: "John Doe", Publisher: "Publisher A"}, nil},
		{"get with invalid bookID", "id=invalid", "invalid", nil, errors.NotFound{Resource: "Book"}},
	}

	store, h, app := initializeHandlersTest(t)

	for i, tc := range tests {
		req := httptest.NewRequest(http.MethodGet, "/book?"+tc.queryParams, nil)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		store.EXPECT().Get(ctx, tc.bookID).Return(tc.resp, tc.err)

		resp, err := h.Get(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestHandler_Create_Invalid_Input_Error(t *testing.T) {
	expErr := errors.Error("test error")

	_, h, app := initializeHandlersTest(t)
	req := httptest.NewRequest(http.MethodPost, "/book", errReader(0))
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	_, err := h.Create(ctx)

	assert.Equal(t, expErr, err)
}

func TestHandler_Create_Invalid_JSON(t *testing.T) {
	input := `{"id":"123","title":"Sample Book","author":"John Doe","publisher":42}`
	expErr := &json.UnmarshalTypeError{
		Value:  "number",
		Type:   reflect.TypeOf(""),
		Offset: 35,
		Struct: "Book",
		Field:  "publisher",
	}

	_, h, app := initializeHandlersTest(t)

	inputReader := strings.NewReader(input)
	req := httptest.NewRequest(http.MethodPost, "/book", inputReader)
	r := request.NewHTTPRequest(req)
	ctx := gofr.NewContext(nil, r, app)

	_, err := h.Create(ctx)

	assert.Equal(t, expErr, err)
}

func TestHandler_Create(t *testing.T) {
	bookJSON := `{"id":"123","title":"Sample Book","author":"John Doe","publisher":"Publisher A"}`
	book := models.Book{ID: "123", Title: "Sample Book", Author: "John Doe", Publisher: "Publisher A"}
	tests := []struct {
		desc string
		resp string
		err  error
	}{
		{"create success", "New book added!", nil},
		{"create fail", "", errors.Error("test error")},
	}

	store, h, app := initializeHandlersTest(t)

	for i, tc := range tests {
		input := strings.NewReader(bookJSON)

		req := httptest.NewRequest(http.MethodPost, "/book", input)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		store.EXPECT().Create(ctx, book).Return(tc.err)

		_, err := h.Create(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestHandler_Delete(t *testing.T) {
	tests := []struct {
		desc  string
		bookID  string
		count int
		resp  interface{}
		err   error
	}{
		{"delete with valid bookID", "123", 1, "1 books deleted!", nil},
		{"delete with invalid bookID", "invalid", 0, nil, errors.NotFound{Resource: "Book"}},
	}

	store, h, app := initializeHandlersTest(t)

	for i, tc := range tests {
		req := httptest.NewRequest(http.MethodDelete, "/book/"+tc.bookID, nil)
		r := request.NewHTTPRequest(req)
		ctx := gofr.NewContext(nil, r, app)

		store.EXPECT().Delete(ctx, tc.bookID).Return(tc.count, tc.err).Times(1)

		resp, err := h.Delete(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

type errReader int

func (errReader) Read([]byte) (n int, err error) {
	return 0, errors.Error("test error")
}