package storage

import (
	"errors"
	"fmt"

	"applicationDesignTest/internals/domain/entities"
)

var (
	ConstraintErr         = errors.New("constraint error")
	WrongOrderDatesErr    = fmt.Errorf("%w: date to less date from", ConstraintErr)
	RoomAlreadyBookingErr = fmt.Errorf("%w: room already booking", ConstraintErr)
)

type Storager interface {
	GetRoomGetter() entities.RoomsStorager
	GetOrderGetter() entities.OrderStorager
}
