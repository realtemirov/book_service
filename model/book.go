package model

type Book struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
	Price       int    `json:"price" db:"price"`
	Author      string `json:"author" db:"author" validate:"required"`
}

type UpdateBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Author      string `json:"author"`
}

type ResponseBook struct {
	Books []Book `json:"books"`
	Count int    `json:"count"`
}

type Request struct {
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
	Search string `json:"search"`
	Sort   string `json:"sort"`
	Asc    bool   `json:"asc"`
}
