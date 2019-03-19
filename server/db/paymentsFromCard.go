package db

import (
	"context"

	"github.com/kadukm/banking_spa/server/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNewPaymentFromCard(payment utils.PaymentFromCardDTO) error {
	res := convertToPaymentFromCard(payment)
	_, err := paymentsFromCard.InsertOne(context.TODO(), res)
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

func GetPaymentsFromCard(sortDTO utils.MongoSortDTO) (res []*utils.PaymentFromCardDTO, err error) {
	sort := convertToSortOption(sortDTO)
	findOptions := options.Find().SetSort(sort)

	cursor, err := paymentsFromCard.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var curPayment paymentFromCard
		err = cursor.Decode(&curPayment)
		if err != nil {
			return nil, err
		}
		curPaymentDTO := curPayment.convertToPaymentFromCardDTO()
		res = append(res, &curPaymentDTO)
	}

	cursor.Close(context.TODO())
	return
}
