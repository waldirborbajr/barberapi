package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"id"`
	Name       string             `json:"name,omitempty" validate:"required"`
	Location   string             `json:"location,omitempty" validate:"required"`
	Title      string             `json:"title,omitempty" validate:"required"`
	Email      string             `json:"email"`
	Password   string             `json:"password"`
	Enable     string             `json:"enable"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	DisabledAt time.Time          `json:"disabled_at"`
}
