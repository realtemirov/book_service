package model

import "time"

type Order struct {
	ID        string     `json:"id" db:"id"`
	BookID    string     `json:"book_id" db:"book_id"`
	UserID    string     `json:"user_id" db:"user_id"`
	Quantity  int        `json:"quantity" db:"quantity"`
	Total     int        `json:"total" db:"total"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type UpdateOrder struct {
	BookId   string     `json:"book_id" db:"book_id"`
	UserId   string     `json:"user_id" db:"user_id"`
	Quantity int        `json:"quantity" db:"quantity"`
	Total    int        `json:"total" db:"total"`
	UpdateAt *time.Time `json:"updated_at" db:"updated_at"`
}

type ResponseOrder struct {
	Orders []Order `json:"orders"`
	Count  int     `json:"count"`
}
