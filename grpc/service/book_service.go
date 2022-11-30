package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/realtemirov/book_service/config"
	"github.com/realtemirov/book_service/genproto/book_service"
	"github.com/realtemirov/book_service/model"
	"github.com/realtemirov/book_service/pkg/logger"
	"github.com/realtemirov/book_service/storage"
)

type bookService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	book_service.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *bookService {
	return &bookService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (b *bookService) CreateBook(ctx context.Context, book *book_service.Book) (*book_service.Book, error) {

	strconv.ParseInt(uuid.New().String(), 10, 32)

	res, err := b.strg.Books().CreateBook(ctx, model.Book{
		ID:          uuid.New().String(),
		Title:       book.GetTitle(),
		Description: book.GetDescription(),
		Author:      book.GetAuthor(),
		Price:       int(book.Price),
	})

	if err != nil {
		return nil, err
	}

	return &book_service.Book{
		Id:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Author:      res.Author,
		Price:       int32(res.Price),
	}, nil
}

func (b *bookService) DeleteBook(ctx context.Context, book *book_service.BookId) (*book_service.Book, error) {
	res, err := b.strg.Books().DeleteBook(ctx, book.BookId)
	if err != nil {
		fmt.Println("Error here book servive: ", err.Error())
		return nil, err
	}
	return &book_service.Book{
		Id:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Price:       int32(res.Price),
		Author:      res.Author,
	}, nil
}

func (b *bookService) GetAllBooks(ctx context.Context, req *book_service.Req) (*book_service.Books, error) {

	filter := fmt.Sprintf("1=1 AND (title ILIKE '%%%s%%') OR (description ILIKE '%%%s%%') OR (author ILIKE '%%%s%%') ", req.Search, req.Search, req.Search)
	sort := fmt.Sprintf(" ORDER BY %s", req.Sort)

	if req.Asc {
		sort += " ASC"
	} else {
		sort += " DESC"
	}

	books, err := b.strg.Books().GetAllBooks(ctx, model.Request{
		Offset: req.Offset,
		Limit:  req.Limit,
		Search: filter,
		Sort:   req.Sort,
		Asc:    req.Asc,
	})

	if err != nil {
		return nil, err
	}

	var bookList []*book_service.Book
	for _, book := range books.Books {
		bookList = append(bookList, &book_service.Book{
			Id:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Price:       int32(book.Price),
		})
	}

	return &book_service.Books{
		Books: bookList,
		Count: int32(books.Count),
	}, nil

}
func (b *bookService) GetBook(ctx context.Context, BookId *book_service.BookId) (*book_service.Book, error) {

	book, err := b.strg.Books().GetBook(ctx, BookId.BookId)
	if err != nil {
		return nil, err
	}
	return &book_service.Book{
		Id:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Price:       int32(book.Price),
		Author:      book.Author,
	}, nil
}

func (b *bookService) UpdateBook(ctx context.Context, book *book_service.NewBook) (*book_service.Book, error) {
	res, err := b.strg.Books().UpdateBook(ctx, book.Id, model.UpdateBook{
		Title:       book.Book.Title,
		Description: book.Book.Description,
		Price:       int(book.Book.Price),
		Author:      book.Book.Author,
	})
	if err != nil {
		return nil, err
	}
	return &book_service.Book{
		Id:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Price:       int32(res.Price),
		Author:      res.Author,
	}, nil
}
