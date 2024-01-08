package usecases

import (
	"context"
	"errors"
	"fmt"

	"applicationDesignTest/internals/domain/entities"
)

var (
	DatabaseErr  = errors.New("database error")
	GetByUserErr = fmt.Errorf("%w: on  call GetByUser")
	PostErr      = fmt.Errorf("%w: on  call Post")
)

type OrdersUseCases interface {
	GetByUser(ctx context.Context, id int64) ([]entities.Order, error)
	OrderPost(ctx context.Context, order entities.Order) (entities.Order, error)
}
