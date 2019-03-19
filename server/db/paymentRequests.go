package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/kadukm/banking_spa/server/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNewPaymentRequest(request utils.PaymentRequestDTO) error {
	res := convertToPaymentRequest(request)
	_, err := paymentRequests.InsertOne(context.TODO(), res)
	return err
}

func GetCompany(companyID string) (res utils.CompanyDTO, err error) {
	filter := bson.M{"_id": companyID}
	companyInstance := company{}
	if err = companies.FindOne(context.TODO(), filter).Decode(&companyInstance); err != nil {
		return
	}
	res = companyInstance.convertToCompanyDTO()
	return
}

func GetPaymentRequests(sortDTO utils.MongoSortDTO) (res []*utils.PaymentRequestDTO, err error) {
	sort := convertToSortOption(sortDTO)
	findOptions := options.Find().SetSort(sort)

	cursor, err := paymentRequests.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var curPayment paymentRequest
		err = cursor.Decode(&curPayment)
		if err != nil {
			return nil, err
		}
		curPaymentDTO := curPayment.convertToPaymentRequestDTO()
		res = append(res, &curPaymentDTO)
	}

	cursor.Close(context.TODO())
	return
}
