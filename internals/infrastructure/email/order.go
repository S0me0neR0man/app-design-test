package email

import (
	"context"

	"applicationDesignTest/internals/domain/entities"
)

type OrderEmailNotify struct {
}

func (o OrderEmailNotify) NotifyAsync(ctx context.Context, order entities.Order) {
	//TODO implement me
	panic("implement me")
}
