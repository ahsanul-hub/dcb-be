package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"app/database"
	"app/dto/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindPaymentMethodBySlug(slug string, defaultValue interface{}) (*model.PaymentMethod, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.GetCollection("redpay", "settings")

	var paymentMethod model.PaymentMethod
	filter := bson.M{"slug": slug}

	err := collection.FindOne(ctx, filter).Decode(&paymentMethod)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &paymentMethod, nil
}

func GetPrice(prefix string, amount float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Collection reference
	collection := database.GetCollection("redpay", "settings")

	// Build slug
	slug := fmt.Sprintf("%s_charging", prefix)

	// Query MongoDB
	var paymentMethod model.PaymentMethod
	filter := bson.M{"slug": slug}
	err := collection.FindOne(ctx, filter).Decode(&paymentMethod)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, fmt.Errorf("charging settings not found for prefix: %s", prefix)
		}
		return 0, err
	}

	// Validate if Denom exists
	if len(paymentMethod.Value.Denom) == 0 {
		return 0, fmt.Errorf("no denominated values available for prefix: %s", prefix)
	}

	// Loop through Denom to find the price for the given amount
	for denom, price := range paymentMethod.Value.Denom {
		// Convert denom (key) to float64 and compare
		denomFloat := convertDenomToFloat(denom)
		if denomFloat == amount {
			return price, nil
		}
	}

	// Return error if no matching denom is found
	return 0, fmt.Errorf("amount %.2f not found in denominated values for prefix: %s", amount, prefix)
}

// Helper function to convert denom key (string) to float64
func convertDenomToFloat(denom string) float64 {
	var denomFloat float64
	fmt.Sscanf(denom, "%f", &denomFloat)
	return denomFloat
}
