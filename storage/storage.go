package storage

import (
	"context"

	"github.com/realtemirov/book_service/model"
)

type StorageI interface {
	Books() BookI
	Orders() OrderI
}

type OrderI interface {
	CreateOrder(ctx context.Context, order model.Order) (*model.Order, error)
	GetOrder(ctx context.Context, id string) (*model.Order, error)
	GetAllOrders(ctx context.Context, req model.Request) (*model.ResponseOrder, error)
	UpdateOrder(ctx context.Context, id string, order *model.UpdateOrder) (*model.Order, error)
	DeleteOrder(ctx context.Context, id string) (*model.Order, error)
}

type BookI interface {
	CreateBook(ctx context.Context, book model.Book) (*model.Book, error)
	GetBook(ctx context.Context, id string) (*model.Book, error)
	GetAllBooks(ctx context.Context, req model.Request) (*model.ResponseBook, error)
	UpdateBook(ctx context.Context, id string, book model.UpdateBook) (*model.Book, error)
	DeleteBook(ctx context.Context, id string) (*model.Book, error)
}
