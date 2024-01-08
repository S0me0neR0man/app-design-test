package entities

import (
	"context"
	"time"
)

type Order struct {
	ID     int64     `json:"id,omitempty"`
	UserID int64     `json:"user_id,omitempty"`
	Rooms  []int64   `json:"rooms,omitempty"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

type OrderStorager interface {
	PostOrder(ctx context.Context, order Order) (Order, error)
	OrdersByUser(ctx context.Context, id int64) ([]Order, error)
}

type OrderNotificatorer interface {
	NotifyAsync(ctx context.Context, order Order)
}
