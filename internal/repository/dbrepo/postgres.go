package dbrepo

import (
	"context"
	"time"

	"github.com.br/Leodf/bookings/internal/model"
)

// InsertReservation inserts a reservation into the database
func (r *postgresDBRepo) InsertReservation(res model.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into reservations 
					(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
					`
	_, err := r.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
