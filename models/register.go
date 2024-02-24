package models

import (
	"example.com/rest-api/db"
)

type Register struct {
	ID      int64
	EventID int64 `binding:"required"`
	UserID  int64 `binding:"required"`
}

func (r *Register) Save() error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)

	return err
}

func GetRegistrationByID(eventId int64, userId int64) (*Register, error) {
	query := "SELECT * FROM registrations WHERE event_id = ? AND user_id = ?"
	row := db.DB.QueryRow(query, eventId, userId)

	var register Register

	err := row.Scan(&register.ID, &register.EventID, &register.UserID)
	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *Register) CancelRegistration() error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.EventID, r.UserID)
	return err
}
