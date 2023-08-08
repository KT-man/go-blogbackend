package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId primitive.ObjectID `json:"id,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}