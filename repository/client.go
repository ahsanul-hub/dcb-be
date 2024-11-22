package repository

import (
	"context"
	"time"

	"app/database"
	"app/dto/model"

	"go.mongodb.org/mongo-driver/bson"
)

func FindClient(clientAppKey, clientAppID string) (*model.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("redpay", "clients")

	var result model.Client
	filter := bson.M{
		"client_appkey": clientAppKey,
		"_id":           clientAppID,
	}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
