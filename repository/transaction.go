package repository

import (
	"app/database"
	"app/dto/model"
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckTransaction(transactionID, appKey, appID string) (*model.Transactions, error) {
	ctx := context.Background()

	collection := database.GetCollection("redpay", "transactions")

	filter := bson.M{
		"merchant_transaction_id": transactionID,
		"client_appkey":           appKey,
		"app_id":                  appID,
	}

	var result model.Transactions

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func CreateOrder(ctx context.Context, input *model.InputPaymentRequest, client *model.Client) (uint, error) {
	// Initialize a new Transactions model
	transaction := model.Transactions{
		ClientAppKey:  input.ClientAppKey,
		StatusCode:    stringToInt(input.StatusCode), // Convert status code to int
		ItemName:      input.Status,
		UserMDN:       input.Mobile,
		Testing:       input.Testing,
		Route:         input.Route,
		PaymentMethod: input.PaymentMethod,
		Currency:      input.Currency,
		Price:         input.Price,
	}

	// Populate additional fields from client data
	transaction.AppID = client.ID                // Mengakses field ID dari struct client
	transaction.MerchantName = client.ClientName // Mengakses field ClientName dari struct client
	transaction.AppKey = client.AppName

	collection := database.GetCollection("redpay", "transactions")
	result, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		return 0, err
	}

	// Mengembalikan ID transaksi yang baru saja dibuat
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return 0, fmt.Errorf("could not convert inserted ID to ObjectID")
	}

	return uint(id.Timestamp().Unix()), nil
}

// Helper function to convert string to int
func stringToInt(value string) int {
	intVal, _ := strconv.Atoi(value)
	return intVal
}
