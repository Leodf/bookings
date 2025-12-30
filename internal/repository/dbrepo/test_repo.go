package dbrepo

import (
	"errors"
	"log"
	"time"

	"github.com.br/Leodf/bookings/internal/model"
)

// InsertReservation inserts a reservation into the database
func (r *testDBRepo) InsertReservation(res model.Reservation) (int, error) {
	// if the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return 0, errors.New("some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (r *testDBRepo) InsertRoomRestriction(rr model.RoomRestrictions) error {
	if rr.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no availability
func (r *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	// set up a test time
	layout := "02/01/2006"
	str := "31/12/2049"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}
	// this is our test to fail the query -- specify 2060-01-01 as start
	testDateToFail, err := time.Parse(layout, "01/01/2060")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return false, errors.New("some error")
	}

	// if the start date is after 2049-12-31, then return false,
	// indicating no availability;
	if start.After(t) {
		return false, nil
	}

	// otherwise, we have availability
	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (r *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]model.Room, error) {
	var rooms []model.Room

	// if the start date is after 2049-12-31, then return empty slice,
	// indicating no rooms are available;
	layout := "02/01/2006"
	str := "31/12/2049"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	testDateToFail, err := time.Parse(layout, "01/01/2060")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return rooms, errors.New("some error")
	}

	if start.After(t) {
		return rooms, nil
	}

	// otherwise, put an entry into the slice, indicating that some room is
	// available for search dates
	room := model.Room{
		ID: 1,
	}
	rooms = append(rooms, room)

	return rooms, nil

}

// GetRoomByID get room by ID
func (r *testDBRepo) GetRoomByID(id int) (model.Room, error) {
	var room model.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil

}

func (r *testDBRepo) GetUserByID(id int) (model.User, error) {
	var u model.User

	return u, nil
}

func (r *testDBRepo) UpdateUser(u model.User) error {
	return nil
}

func (r *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

func (r *testDBRepo) AllReservations() ([]model.Reservation, error) {
	var reservations []model.Reservation

	return reservations, nil
}

func (r *testDBRepo) AllNewReservations() ([]model.Reservation, error) {
	var reservations []model.Reservation

	return reservations, nil
}

func (r *testDBRepo) GetReservationByID(id int) (model.Reservation, error) {

	var res model.Reservation

	return res, nil
}

func (r *testDBRepo) UpdateReservation(rm model.Reservation) error {

	return nil
}

func (r *testDBRepo) DeleteReservation(id int) error {

	return nil
}

func (r *testDBRepo) UpdateProcessedForReservation(id, processed int) error {

	return nil
}
