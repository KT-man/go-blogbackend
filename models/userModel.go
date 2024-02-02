package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId primitive.ObjectID `json:"id,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

type NewUser struct {
	Username string
	Password string `json:"password" validate:"required,min=6"`
}