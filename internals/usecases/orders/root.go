package orders

import (
	"context"
	"errors"
	"log"

	"applicationDesignTest/internals/domain/entities"
	"applicationDesignTest/internals/domain/usecases"
)

type UseCasesImpl struct {
	orderStorage entities.OrderStorager
	roomStorage  entities.RoomsStorager
	logger       *log.Logger
}

func NewOrderUseCases(orderStorage entities.OrderStorager, roomStorage entities.RoomsStorager, logger *log.Logger) *UseCasesImpl {
	return &UseCasesImpl{
		orderStorage: orderStorage,
		roomStorage:  roomStorage,
		logger:       logger,
	}
}

func (u *UseCasesImpl) GetByUser(ctx context.Context, id int64) ([]entities.Order, error) {
	orders, err := u.orderStorage.OrdersByUser(ctx, id)
	if err != nil {
		return nil, errors.Join(usecases.GetByUserErr, err)
	}

	return orders, nil
}

func (u *UseCasesImpl) OrderPost(ctx context.Context, order entities.Order) (entities.Order, error) {
	posted, err := u.orderStorage.PostOrder(ctx, order)
	if err != nil {
		return order, errors.Join(usecases.PostErr, err)
	}

	return posted, nil
}
