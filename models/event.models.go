package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int64      `json:"user_id"`
}

func (e *Event) TableName() string {
	return "events"
}

// save event to database later
var events = []Event{}

func (e *Event) Save() error {
	query :=
		"INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	e.ID, err = result.LastInsertId()
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, date_time, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events = []Event{}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT id, name, description, location, date_time, user_id FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e *Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, date_time = ?, user_id = ? 
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}