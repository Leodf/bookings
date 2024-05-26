package repository

import "github.com.br/Leodf/bookings/internal/model"

type DatabaseRepo interface {
	InsertReservation(res model.Reservation) (int, error)
	InsertRoomRestriction(rr model.RoomRestrictions) error
}
