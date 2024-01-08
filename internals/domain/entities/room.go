package entities

import (
	"context"
	"time"
)

type RoomCategory int

func (c RoomCategory) String() string {
	switch c {
	case Economy:
		return EconomyText
	case Standart:
		return StandartText
	case Lux:
		return LuxText
	}

	return "unknown"
}

const (
	Economy  RoomCategory = 0x01
	Standart RoomCategory = 0x02
	Lux      RoomCategory = 0x04

	EconomyText  = "econom"
	StandartText = "standart"
	LuxText      = "lux"
)

type Room struct {
	ID       int64
	Category RoomCategory
}

type RoomFindParameters struct {
	RoomCategory
	From time.Time
	To   time.Time
}

type RoomsStorager interface {
	ByID(ctx context.Context, id int64)
	Find(ctx context.Context, params RoomFindParameters) ([]Room, error)
	Book(context.Context, int64) error
}
