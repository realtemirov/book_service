package postgres

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/realtemirov/book_service/config"
	"github.com/realtemirov/book_service/model"
	"github.com/realtemirov/book_service/storage"
	"github.com/stretchr/testify/require"
)

var repo storage.BookI

func TestMain(m *testing.M) {
	var err error

	var cnf = config.Config{
		PostgresHost:     "localhost",
		PostgresPort:     7050,
		PostgresUser:     "postgres",
		PostgresPassword: "postgres",
		PostgresDatabase: "test_db",
		PostgresSSLMode:  "disable",

		PostgresMaxConnections:  30,
		PostgresConnMaxIdleTime: 10,
	}

	storageI, err := NewPostgres(cnf)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	repo = storageI.Books()
	repo2 = storageI.Orders()
	os.Exit(m.Run())
}
func createRandomBook(t *testing.T) *model.Book {
	args := model.Book{
		ID:          uuid.New().String(),
		Title:       uuid.New().String(),
		Description: uuid.New().String(),
		Author:      uuid.New().String(),
		Price:       rand.Int(),
	}

	res, err := repo.CreateBook(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.Title, args.Title)
	require.Equal(t, res.Description, args.Description)
	require.Equal(t, res.Author, args.Author)
	require.Equal(t, res.Price, args.Price)
	require.NotZero(t, res.ID)

	return res
}
func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

/*func TestGetBookError(t *testing.T) {

	book1 := createRandomBook(t)
	book2, err := repo.GetBook(context.Background(), "3")

	require.Error(t, err)
	require.NotEqual(t, book1.ID, book2.ID)
	require.NotEqual(t, book1.Title, book2.Title)
	require.NotEqual(t, book1.Description, book2.Description)
	require.NotEqual(t, book1.Author, book2.Author)
	require.NotEqual(t, book1.Price, book2.Price)
}*/

func TestGetBook(t *testing.T) {

	acc1 := createRandomBook(t)
	acc2, err := repo.GetBook(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Title, acc2.Title)
	require.Equal(t, acc1.Description, acc2.Description)
	require.Equal(t, acc1.Author, acc2.Author)
	require.Equal(t, acc1.Price, acc2.Price)
}

func TestGetAllBooks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBook(t)
	}

	res, err := repo.GetAllBooks(context.Background(), model.Request{
		Offset: 0,
		Limit:  10,
		Search: "1=1",
		Sort:   "id ASC",
		Asc:    false,
	})
	require.NoError(t, err)

	for _, b := range res.Books {
		require.NotEmpty(t, b)
	}
}

func TestUpdateBook(t *testing.T) {
	book1 := createRandomBook(t)

	upt := model.UpdateBook{
		Title:       book1.Title,
		Description: "Update",
		Author:      book1.Author,
		Price:       book1.Price,
	}
	book2, err := repo.UpdateBook(context.Background(), book1.ID, upt)

	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, upt.Description, book2.Description)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Price, book2.Price)
}
func TestDeleteBook(t *testing.T) {
	book1 := createRandomBook(t)
	book2, err := repo.DeleteBook(context.Background(), book1.ID)
	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Description, book2.Description)
	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Price, book2.Price)

	require.NoError(t, err)

	book3, err := repo.GetBook(context.Background(), book1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book3)
}
