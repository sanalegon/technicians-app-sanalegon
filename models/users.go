package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	Name      string             `bson: "name" json: "name, omitempty"`
	Lastname  string             `bson: "lastname" json: "lastname, omitempty"`
	Birthday  time.Time          `bson: "birhtday" json: "birhtday, omitempty"`
	Email     string             `bson: "email json: "email"`                   // Aqui siempre devolvera el email
	Password  string             `bson: "password" json: "password, omitempty"` // Siempre omitempty porque nunca puedo regresar un pass por el navegador
	Avatar    string             `bson: "avatar" json: "avatar, omitempty"`
	Banner    string             `bson: "banner" json: "banner, omitempty"`
	Biography string             `bson: "biography" json: "biography, omitempty"`
	Website   string             `bson: "website" json: "website, omitempty"`
	Location  string             `bson: "location" json: "location, omitempty"`
}
