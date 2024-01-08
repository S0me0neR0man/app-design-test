package simple

import (
	"context"
	"sync"

	"applicationDesignTest/internals/domain/entities"
	"applicationDesignTest/utils"
)

type Database struct {
	nextOrderId int64

	orders   []entities.Order
	ordersMu sync.RWMutex

	rooms   []entities.Room
	roomsMu sync.RWMutex
}

func NewOnSliceDatabase() *Database {
	return &Database{
		orders:      make([]entities.Order, 0),
		rooms:       make([]entities.Room, 0),
		nextOrderId: 100,
	}
}

// *** RoomStorager ***

func (d *Database) ByID(ctx context.Context, id int64) {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Find(ctx context.Context, params entities.RoomFindParameters) ([]entities.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Book(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

// *** OrderStorager ***

// PostOrder
// TODO: add checks
func (d *Database) PostOrder(ctx context.Context, order entities.Order) (entities.Order, error) {
	d.ordersMu.Lock()
	defer d.ordersMu.Unlock()

	select {
	case <-ctx.Done():
		return order, utils.ContextCanceledErr
	default:
		order.ID = d.nextOrderId
		d.nextOrderId++
		d.orders = append(d.orders, order)
	}

	return order, nil
}

func (d *Database) OrdersByUser(ctx context.Context, id int64) ([]entities.Order, error) {
	d.ordersMu.RLock()
	defer d.ordersMu.RUnlock()

	orders := make([]entities.Order, 0)

	for _, order := range d.orders {
		select {
		case <-ctx.Done():
			return nil, utils.ContextCanceledErr
		default:
			if order.UserID == id {
				orders = append(orders, order)
			}
		}
	}

	return orders, nil
}
