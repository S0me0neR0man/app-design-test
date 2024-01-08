package controllers

import (
	"context"

	"applicationDesignTest/internals/domain/entities"
	"applicationDesignTest/internals/domain/usecases"
)

type OrdersController struct {
	ordersUseCases usecases.OrdersUseCases
}

func NewOrdersController(ordersUseCases usecases.OrdersUseCases) *OrdersController {
	return &OrdersController{
		ordersUseCases: ordersUseCases,
	}
}

func (o *OrdersController) GetByUser(ctx context.Context, id int64) ([]entities.Order, error) {
	return o.ordersUseCases.GetByUser(ctx, id)
}

func (o *OrdersController) OrderPost(ctx context.Context, order entities.Order) (entities.Order, error) {
	return o.ordersUseCases.OrderPost(ctx, order)
}
