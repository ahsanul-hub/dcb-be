package database

import (
	"app/config"
	"app/dto/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MongoClient *mongo.Client

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_PASSWORD", ""), config.Config("DB_HOST", "localhost"), config.Config("DB_PORT", ""), config.Config("DB_NAME", ""))
	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Product{}, &model.User{})
	fmt.Println("Database Migrated")
}

func SetupMongoDB() {
	uri := config.Config("MONGODB_URI", "")
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("failed to connect to MongoDB")
	}

	log.Println("Connected to MongoDB")
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return MongoClient.Database(databaseName).Collection(collectionName)
}
