package user

import (
	"time"
)

type User struct {
	Id         int       `json:"id" :"id"`
	Name       string    `json:"user_name" binding:"required" :"name"`
	DayOfBirth string    `json:"day_of_birth" binding:"required" : "day_of_birth"`
	Phone      string    `json:"phone" binding:"required" :"phone"`
	Email      string    `json:"email" binding:"required" :"email"`
	UserName   string    `json:"user_name" binding:"required" :"user_name"`
	PassWord   string    `json:"pass_word" binding:"required" :"pass_word"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
