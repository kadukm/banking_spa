package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var paymentsFromCard mongo.Collection
var paymentRequests mongo.Collection
var companies mongo.Collection
var products mongo.Collection
var users mongo.Collection

func init() {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("banking-spa")
	paymentsFromCard = *db.Collection("payments-from-card")
	paymentRequests = *db.Collection("payment-requests")
	companies = *db.Collection("companies")
	products = *db.Collection("products")
	users = *db.Collection("users")
}
