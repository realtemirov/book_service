package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/book_service/config"
	"github.com/realtemirov/book_service/storage"
)

type Store struct {
	db *sqlx.DB

	bookRepo  storage.BookI
	orderRepo storage.OrderI
}

func NewPostgres(cnf config.Config) (storage.StorageI, error) {
	psqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.PostgresHost,
		cnf.PostgresPort,
		cnf.PostgresUser,
		cnf.PostgresPassword,
		cnf.PostgresDatabase,
	)
	db, err := sqlx.Open("postgres", psqlConnection)
	if err != nil {
		log.Fatalf("cannot connect to postgres: %v", err)
	}

	db.SetConnMaxIdleTime(time.Duration(time.Duration(cnf.PostgresConnMaxIdleTime).Minutes()))
	db.SetMaxOpenConns(cnf.PostgresMaxConnections)

	if err = db.Ping(); err != nil {
		log.Fatalf("cannot connect to postgres: %s", err.Error())
	}

	return &Store{
		db: db,
	}, nil

}

func (s *Store) Books() storage.BookI {
	if s.bookRepo == nil {
		s.bookRepo = NewBookRepo(s.db)
	}

	return s.bookRepo
}

func (s *Store) Orders() storage.OrderI {
	if s.orderRepo == nil {
		s.orderRepo = NewOrderRepo(s.db)
	}

	return s.orderRepo
}
