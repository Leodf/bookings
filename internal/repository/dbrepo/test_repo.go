package dbrepo

import (
	"time"

	"github.com.br/Leodf/bookings/internal/model"
)

// InsertReservation inserts a reservation into the database
func (r *testDBRepo) InsertReservation(res model.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (r *testDBRepo) InsertRoomRestriction(rr model.RoomRestrictions) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no availability
func (r *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (r *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]model.Room, error) {
	var rooms []model.Room
	return rooms, nil

}

// GetRoomByID get room by ID
func (r *testDBRepo) GetRoomByID(id int) (model.Room, error) {
	var room model.Room
	return room, nil

}
