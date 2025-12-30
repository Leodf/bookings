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
	GetRoomByID(id int) (model.Room, error)
	GetUserByID(id int) (model.User, error)
	UpdateUser(u model.User) error
	Authenticate(email, testPassword string) (int, string, error)
	AllReservations() ([]model.Reservation, error)
	AllNewReservations() ([]model.Reservation, error)
	GetReservationByID(id int) (model.Reservation, error)
}
