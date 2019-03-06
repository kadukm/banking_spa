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

func AddNewPaymentFromCard(payment utils.PaymentFromCardDTO) error {
	res := convertToPaymentFromCard(payment)
	_, err := paymentsFromCard.InsertOne(context.TODO(), res)
	return err
}

func AddNewPaymentRequest(request utils.PaymentRequestDTO) error {
	res := convertToPaymentRequest(request)
	_, err := paymentRequests.InsertOne(context.TODO(), res)
	return err
}

func PatchPaymentFromCard(patch utils.PatchPaymentFromCardDTO, paymentID string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(paymentID)
	if err != nil {
		return
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": patch}
	_, err = paymentsFromCard.UpdateOne(context.TODO(), filter, update)
	return
}
