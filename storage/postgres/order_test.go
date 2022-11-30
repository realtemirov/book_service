package postgres

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/realtemirov/book_service/model"
	"github.com/realtemirov/book_service/storage"
	"github.com/stretchr/testify/require"
)

var repo2 storage.OrderI

func createRandomOrder(t *testing.T) *model.Order {
	createdAt := time.Now()
	args := model.Order{
		ID:        uuid.New().String(),
		BookID:    uuid.New().String(),
		UserID:    uuid.New().String(),
		Quantity:  rand.Int(),
		Total:     rand.Int(),
		CreatedAt: &createdAt,
		UpdatedAt: nil,
	}
	res, err := repo2.CreateOrder(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.ID, args.ID)
	require.Equal(t, res.BookID, args.BookID)
	require.Equal(t, res.Quantity, args.Quantity)
	require.Equal(t, res.CreatedAt.Second(), args.CreatedAt.Second())
	require.Equal(t, res.Total, args.Total)
	require.Equal(t, res.UserID, args.UserID)
	require.Equal(t, res.UpdatedAt, args.UpdatedAt)
	require.NotZero(t, res.ID)

	return res
}
func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

/*func TestGetOrderError(t *testing.T) {

	Order1 := createRandomOrder(t)
	Order2, err := repo.GetOrder(context.Background(), "3")

	require.Error(t, err)
	require.NotEqual(t, Order1.ID, Order2.ID)
	require.NotEqual(t, Order1.Title, Order2.Title)
	require.NotEqual(t, Order1.Description, Order2.Description)
	require.NotEqual(t, Order1.Author, Order2.Author)
	require.NotEqual(t, Order1.Price, Order2.Price)
}*/

func TestGetOrder(t *testing.T) {

	acc1 := createRandomOrder(t)
	acc2, err := repo2.GetOrder(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.BookID, acc2.BookID)
	require.Equal(t, acc1.Quantity, acc2.Quantity)
	require.Equal(t, acc1.CreatedAt, acc2.CreatedAt)
	require.Equal(t, acc1.Total, acc2.Total)
	require.Equal(t, acc1.UserID, acc2.UserID)
	require.Equal(t, acc1.UpdatedAt, acc2.UpdatedAt)
	require.NotZero(t, acc2.ID)

}

func TestGetAllOrder(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	res, err := repo2.GetAllOrders(context.Background(), model.Request{
		Offset: 0,
		Limit:  10,
		Search: "1=1",
		Sort:   " ORDER BY id DESC ",
		Asc:    false,
	})
	require.NoError(t, err)

	for _, b := range res.Orders {
		require.NotEmpty(t, b)
	}
	require.Equal(t, len(res.Orders), res.Count)
}

func TestUpdateOrder(t *testing.T) {
	o1 := createRandomOrder(t)

	h := time.Now().Add(time.Hour * 96)

	upt := &model.UpdateOrder{
		BookId:   "nulllll",
		UserId:   o1.UserID,
		Quantity: o1.Quantity * 10,
		Total:    o1.Total * 0,
		UpdateAt: &h,
	}
	o2, err := repo2.UpdateOrder(context.Background(), o1.ID, upt)

	require.NoError(t, err)
	require.NotEmpty(t, o2)
	require.Equal(t, o1.ID, o2.ID)
	require.Equal(t, upt.BookId, o2.BookID)
	require.Equal(t, upt.UpdateAt.Second(), o2.UpdatedAt.Second())
	require.Equal(t, upt.UserId, o2.UserID)
	require.Equal(t, upt.Quantity, o2.Quantity)
	require.Equal(t, upt.Total, o2.Total)

}
func TestDeleteOrder(t *testing.T) {
	o1 := createRandomOrder(t)
	o2, err := repo2.DeleteOrder(context.Background(), o1.ID)
	require.Equal(t, o1.ID, o2.ID)
	require.Equal(t, o1.BookID, o2.BookID)
	require.Equal(t, o1.UpdatedAt, o2.UpdatedAt)
	require.Equal(t, o1.UserID, o2.UserID)
	require.Equal(t, o1.Quantity, o2.Quantity)
	require.Equal(t, o1.Total, o2.Total)
	require.NoError(t, err)

	o3, err := repo2.GetOrder(context.Background(), o1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, o3)
}
