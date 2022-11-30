package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/book_service/model"
	"github.com/spf13/cast"
)

type bookRepo struct {
	db *sqlx.DB
}

var (
	bookTable = "books"
)

func NewBookRepo(db *sqlx.DB) *bookRepo {
	return &bookRepo{db}
}

// CreateBook creates a new book
func (b *bookRepo) CreateBook(ctx context.Context, book model.Book) (*model.Book, error) {

	var res model.Book
	query := fmt.Sprintf("INSERT INTO %s (id, title, author, description, price) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, author, description, price", bookTable)
	row := b.db.QueryRowContext(ctx, query, book.ID, book.Title, book.Author, book.Description, book.Price)
	err := row.Scan(&res.ID, &res.Title, &res.Author, &res.Description, &res.Price)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetBook gets a book by id
func (b *bookRepo) GetBook(ctx context.Context, id string) (*model.Book, error) {

	query := fmt.Sprintf("SELECT id,title,author,description,price FROM %s WHERE id = $1", bookTable)
	var book model.Book
	row := b.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// GetAllBooks gets all books
func (b *bookRepo) GetAllBooks(ctx context.Context, req model.Request) (*model.ResponseBook, error) {

	var books model.ResponseBook

	query := fmt.Sprintf("SELECT id,title,author,description,price FROM %s WHERE %s ORDER BY %v OFFSET %v LIMIT %v", bookTable, req.Search, req.Sort, cast.ToString(req.Offset), cast.ToString(req.Limit))
	fmt.Println(query)
	rows, err := b.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price)
		if err != nil {
			return nil, err
		}
		books.Books = append(books.Books, book)
	}
	books.Count = len(books.Books)
	return &books, nil
}

// UpdateBook updates a book by id
func (b *bookRepo) UpdateBook(ctx context.Context, id string, book model.UpdateBook) (*model.Book, error) {

	query := fmt.Sprintf("UPDATE %s SET title = $1, author = $2, description = $3, price = $4 WHERE id = $5 RETURNING title, author, description, price", bookTable)
	err := b.db.QueryRowContext(ctx, query, book.Title, book.Author, book.Description, book.Price, id).Scan(&book.Title, &book.Author, &book.Description, &book.Price)

	if err != nil {
		return nil, err
	}
	return &model.Book{
		ID:          id,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		Price:       book.Price,
	}, nil
}

// DeleteBook deletes a book by id
func (b *bookRepo) DeleteBook(ctx context.Context, id string) (*model.Book, error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id,title, author, description, price", bookTable)
	var book model.Book
	row := b.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
