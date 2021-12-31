package db

import (
	"context"
	"time"

	"github.com/sanalegon/technicians-app-sanalegon/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* CheckIfUserExists receives an email as parameter and checks if it is already in the database */
func CheckIfUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("technicians-app") // TODO: Save these variables as constants
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
