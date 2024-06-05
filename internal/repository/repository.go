package repository

import (
	"time"

	"github.com.br/Leodf/bookings/internal/model"
)

type DatabaseRepo interface {
	InsertReservation(res model.Reservation) (int, error)
	InsertRoomRestriction(rr model.RoomRestrictions) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]model.Room, error)
}
