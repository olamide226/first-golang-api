package models

import (
	"example.com/rest-api/utils"

	"example.com/rest-api/db"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type UserLogin struct {
	ID	   int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Save() error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)

	if err != nil {
		return err
	}
	u.ID, err = result.LastInsertId()
	return err
}

func (u *UserLogin) Authenticate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return err
	}
	err = utils.CheckPasswordHash(u.Password, hashedPassword)
	return err
}
