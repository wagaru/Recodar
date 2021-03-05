package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email         string             `json:"email"`
	Name          string             `json:"name"`
	Picture       string             `json:"picture"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	BindingSource string             `json:"binding_source" bson:"binding_source"`
}
