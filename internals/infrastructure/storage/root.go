package storage

import "applicationDesignTest/internals/domain/entities"

type Storager interface {
	GetRoomGetter() entities.RoomsStorager
	GetOrderGetter() entities.OrderStorager
}
