package simple

import (
	"context"
	"sync"
	"time"

	"applicationDesignTest/internals/domain/entities"
	"applicationDesignTest/internals/infrastructure/storage"
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

func (d *Database) PostOrder(ctx context.Context, order entities.Order) (entities.Order, error) {
	d.ordersMu.Lock()
	defer d.ordersMu.Unlock()

	select {
	case <-ctx.Done():
		return order, utils.ContextCanceledErr
	default:
		if err := d.isOrderOk(ctx, order); err != nil {
			return order, err
		}
		order.ID = d.nextOrderId
		d.nextOrderId++
		d.orders = append(d.orders, order)
	}

	return order, nil
}

// TODO: refactor to Rooms usecases
func (d *Database) isOrderOk(ctx context.Context, order entities.Order) error {
	if order.To.Before(order.From) {
		return storage.WrongOrderDatesErr
	}
	for _, o := range d.orders {
		select {
		case <-ctx.Done():
			return utils.ContextCanceledErr
		default:
		}
		if isOverlap(o.From, o.To, order.From, order.To) {
			m := make(map[int64]bool)
			for _, room := range order.Rooms {
				m[room] = true
			}
			for _, room := range o.Rooms {
				if m[room] {
					return storage.RoomAlreadyBookingErr
				}
			}
		}
	}
	return nil
}

func isOverlap(start1, end1, start2, end2 time.Time) bool {
	return start1.Before(end2) && end1.After(start2)
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
