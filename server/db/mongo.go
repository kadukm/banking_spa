package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kadukm/banking_spa/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var paymentsFromCard mongo.Collection
var paymentRequests mongo.Collection

func init() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("banking-spa-testing")
	paymentsFromCard = *db.Collection("payments-from-card")
	paymentRequests = *db.Collection("payment-requests")
}

func AddNewPaymentFromCard(payment utils.PaymentFromCardDTO) {
	paymentsFromCard.InsertOne(context.TODO(), payment)
}

func AddNewPaymentRequest(request utils.PaymentRequestDTO) {
	paymentRequests.InsertOne(context.TODO(), request)
}

func PatchPaymentFromCard(patch utils.PatchPaymentFromCardDTO, paymentID string) {
	objectID, _ := primitive.ObjectIDFromHex(paymentID)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": patch}
	paymentsFromCard.UpdateOne(context.TODO(), filter, update)
}
