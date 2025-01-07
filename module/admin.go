package module

import (
	"context"
	"time"

	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAdminByUsername(collection *mongo.Collection, username string) (*model.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admin model.Admin
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
