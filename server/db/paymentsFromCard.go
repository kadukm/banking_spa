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
	_, err := paymentsFromCard.InsertOne(context.Background(), res)
	return err
}

func PatchPaymentFromCard(patch utils.PatchPaymentFromCardDTO, paymentID string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(paymentID)
	if err != nil {
		return
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": patch}
	_, err = paymentsFromCard.UpdateOne(context.Background(), filter, update)
	return
}

func GetPaymentsFromCardSorted(sortDTO utils.MongoSortDTO) (res []*utils.PaymentFromCardDTO, err error) {
	sort := convertToSortOption(sortDTO)
	findOptions := options.Find().SetSort(sort)

	cursor, err := paymentsFromCard.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var curPayment paymentFromCard
		err = cursor.Decode(&curPayment)
		if err != nil {
			return nil, err
		}
		curPaymentDTO := curPayment.convertToPaymentFromCardDTO()
		res = append(res, &curPaymentDTO)
	}

	cursor.Close(context.Background())
	return
}

func GetPaymentsFromCard(filterDTO utils.MongoFilterDTO) (res []*utils.PaymentFromCardDTO, err error) {
	filter := convertToFilter(filterDTO)
	cursor, err := paymentsFromCard.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var curPayment paymentFromCard
		err = cursor.Decode(&curPayment)
		if err != nil {
			return nil, err
		}
		curPaymentDTO := curPayment.convertToPaymentFromCardDTO()
		res = append(res, &curPaymentDTO)
	}

	cursor.Close(context.Background())
	return
}
