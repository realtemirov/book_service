package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/book_service/model"
	"github.com/spf13/cast"
)

type orderRepo struct {
	db *sqlx.DB
}

var (
	orderTable = "orders"
)

func NewOrderRepo(db *sqlx.DB) *orderRepo {
	return &orderRepo{db}
}
func (o *orderRepo) CreateOrder(ctx context.Context, order model.Order) (*model.Order, error) {
	query := fmt.Sprintf("INSERT INTO %s (id,book_id,user_id, quantity,total,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,book_id,user_id,quantity,total,created_at,updated_at", orderTable)
	var res model.Order
	row := o.db.QueryRowContext(ctx, query, order.ID, order.BookID, order.UserID, order.Quantity, order.Total, order.CreatedAt, order.UpdatedAt)
	err := row.Scan(&res.ID, &res.BookID, &res.UserID, &res.Quantity, &res.Total, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		fmt.Println("error", err.Error())
		return nil, err
	}
	// fmt.Println("siccess")
	return &res, nil
}
func (o *orderRepo) GetOrder(ctx context.Context, id string) (*model.Order, error) {

	query := fmt.Sprintf("SELECT id,book_id,user_id,quantity,total,created_at,updated_at FROM %s WHERE id = $1", orderTable)
	var order model.Order
	row := o.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&order.ID, &order.BookID, &order.UserID, &order.Quantity, &order.Total, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
func (o *orderRepo) GetAllOrders(ctx context.Context, req model.Request) (*model.ResponseOrder, error) {

	var orders model.ResponseOrder

	query := fmt.Sprintf("SELECT id,book_id,user_id,quantity,total,created_at,updated_at FROM %s", orderTable)
	query = query + " WHERE " + req.Search + req.Sort + " OFFSET " + cast.ToString(req.Offset) + " LIMIT " + cast.ToString(req.Limit)

	ress, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for ress.Next() {
		var order model.Order
		err := ress.Scan(&order.ID, &order.BookID, &order.UserID, &order.Quantity, &order.Total, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders.Orders = append(orders.Orders, order)
	}
	orders.Count = len(orders.Orders)
	return &orders, nil
}
func (o *orderRepo) UpdateOrder(ctx context.Context, id string, order *model.UpdateOrder) (*model.Order, error) {

	var res model.Order
	query := fmt.Sprintf("UPDATE %s SET book_id = $1, user_id = $2, quantity = $3, total = $4, updated_at = $5 WHERE id = $6 RETURNING id, book_id, user_id, quantity, total, created_at, updated_at", orderTable)

	row := o.db.QueryRowContext(ctx, query, order.BookId, order.UserId, order.Quantity, order.Total, order.UpdateAt, id)
	err := row.Scan(&res.ID, &res.BookID, &res.UserID, &res.Quantity, &res.Total, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
func (o *orderRepo) DeleteOrder(ctx context.Context, id string) (*model.Order, error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id, book_id, user_id, quantity, total, created_at, updated_at", orderTable)
	var order model.Order
	row := o.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&order.ID, &order.BookID, &order.UserID, &order.Quantity, &order.Total, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
