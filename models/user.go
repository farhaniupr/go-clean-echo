package models

import "time"

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (u *User) SetTime() {
	u.Created_at = time.Now()
	u.Updated_at = time.Now()
}
