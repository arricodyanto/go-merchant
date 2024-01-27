package entity

import "time"

type Customer struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Contact    string    `json:"contact"`
	IsLoggedIn bool      `json:"is_logged_in"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
