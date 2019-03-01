package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/kadukm/banking_spa/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var paymentsFromCard mongo.Collection
var paymentRequests mongo.Collection

func init() {
	ctx := context.TODO()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("banking-spa-testing")
	paymentsFromCard = *db.Collection("payments-from-card")
	paymentRequests = *db.Collection("payment-requests")
}

func AddNewPaymentFromCard(payment utils.PaymentFromCard) {
	payment.ID = generatorForPaymentsFromCard()
	paymentsFromCard.InsertOne(context.TODO(), payment)
}

func AddNewPaymentRequest(request utils.PaymentRequest) {
	request.ID = generatorForPaymentsRequests()
	paymentRequests.InsertOne(context.TODO(), request)
}

func PatchPaymentFromCard(patch utils.PaymentFromCardPatch, paymentId uint64) {
	filter := bson.D{{"id", paymentId}}
	var payment utils.PaymentFromCard
	paymentsFromCard.FindOne(context.TODO(), filter).Decode(&payment)
	update := bson.D{
		{"$set", bson.D{{"dangerous", true}}},
	}
	paymentsFromCard.UpdateOne(context.TODO(), filter, update)
}

func idGenerator() func() uint64 {
	var i uint64
	return func() uint64 {
		i++
		return i
	}
}

var generatorForPaymentsFromCard = idGenerator()
var generatorForPaymentsRequests = idGenerator()
