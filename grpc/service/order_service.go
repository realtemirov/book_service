package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/realtemirov/book_service/config"
	"github.com/realtemirov/book_service/genproto/order_service"
	"github.com/realtemirov/book_service/model"
	"github.com/realtemirov/book_service/pkg/logger"
	"github.com/realtemirov/book_service/storage"
	"github.com/spf13/cast"
)

type orderService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *orderService {
	return &orderService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (o *orderService) GetOrder(ctx context.Context, id *order_service.OrderId) (*order_service.Order, error) {

	res, err := o.strg.Orders().GetOrder(ctx, id.OrderId)
	if err != nil {
		return nil, err
	}

	return &order_service.Order{
		Id:        res.ID,
		BookId:    res.BookID,
		UserId:    res.UserID,
		Quantity:  strconv.Itoa(res.Quantity),
		Total:     strconv.Itoa(res.Total),
		CreatedAt: cast.ToString(res.CreatedAt),
		UpdatedAt: cast.ToString(res.UpdatedAt),
	}, nil
}

func (o *orderService) GetAllOrders(ctx context.Context, req *order_service.Req) (*order_service.Orders, error) {
	filter := fmt.Sprintf("1=1 AND (book_id ILIKE '%%%s%%') OR (user_id ILIKE '%%%s%%') ", req.Search, req.Search)
	sort := fmt.Sprintf(" ORDER BY %s", req.Sort)
	if req.Asc {
		sort += " ASC"
	} else {
		sort += " DESC"
	}

	orders, err := o.strg.Orders().GetAllOrders(ctx, model.Request{
		Offset: req.Offset,
		Limit:  req.Limit,
		Search: filter,
		Sort:   sort,
		Asc:    req.Asc,
	})

	if err != nil {
		return nil, err
	}

	var orderList []*order_service.Order
	for _, order := range orders.Orders {
		or := &order_service.Order{
			Id:        order.ID,
			BookId:    order.BookID,
			UserId:    order.UserID,
			Quantity:  strconv.Itoa(order.Quantity),
			Total:     strconv.Itoa(order.Total),
			CreatedAt: order.CreatedAt.String(),
		}
		if order.UpdatedAt != nil {
			or.UpdatedAt = order.UpdatedAt.String()
		}
		orderList = append(orderList, or)

	}

	return &order_service.Orders{
		Orders: orderList,
	}, nil
}

func (o *orderService) UpdateBook(ctx context.Context, order *order_service.NewOrder) (*order_service.Order, error) {
	upt := time.Now()
	res, err := o.strg.Orders().UpdateOrder(ctx, order.Id, &model.UpdateOrder{
		BookId:   order.Order.BookId,
		UserId:   order.Order.UserId,
		Quantity: cast.ToInt(order.Order.Quantity),
		Total:    cast.ToInt(order.Order.Total),
		UpdateAt: &upt,
	})

	if err != nil {
		return nil, err
	}

	return &order_service.Order{
		Id:        res.ID,
		BookId:    res.BookID,
		UserId:    res.UserID,
		Quantity:  strconv.Itoa(res.Quantity),
		Total:     strconv.Itoa(res.Total),
		CreatedAt: cast.ToString(res.CreatedAt),
		UpdatedAt: cast.ToString(res.UpdatedAt),
	}, nil
}

func (o *orderService) CreateOrder(ctx context.Context, order *order_service.Order) (*order_service.Order, error) {
	createdAt := time.Now()
	res, err := o.strg.Orders().CreateOrder(ctx, model.Order{
		ID:        uuid.NewString(),
		BookID:    order.BookId,
		UserID:    order.UpdatedAt,
		Quantity:  cast.ToInt(order.Quantity),
		Total:     cast.ToInt(order.Total),
		CreatedAt: &createdAt,
	})

	if err != nil {
		return nil, err
	}
	if res.UpdatedAt == nil {
		createdAt = cast.ToTime("2020-01-01 00:00:00")
		res.UpdatedAt = &createdAt
	}
	return &order_service.Order{
		Id:        res.ID,
		BookId:    res.BookID,
		UserId:    res.UserID,
		Quantity:  strconv.Itoa(res.Quantity),
		Total:     strconv.Itoa(res.Total),
		CreatedAt: res.CreatedAt.String(),
		UpdatedAt: res.UpdatedAt.String(),
	}, nil
}

func (o *orderService) DeleteOrder(ctx context.Context, id *order_service.OrderId) (*order_service.Order, error) {
	res, err := o.strg.Orders().DeleteOrder(ctx, id.OrderId)
	if err != nil {
		return nil, err
	}

	return &order_service.Order{
		Id:        res.ID,
		BookId:    res.BookID,
		UserId:    res.UserID,
		Quantity:  strconv.Itoa(res.Quantity),
		Total:     strconv.Itoa(res.Total),
		CreatedAt: cast.ToString(res.CreatedAt),
		UpdatedAt: cast.ToString(res.UpdatedAt),
	}, nil
}
