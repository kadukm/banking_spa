package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/kadukm/banking_spa/server/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddNewPaymentRequest(request utils.PaymentRequestDTO) error {
	res := convertToPaymentRequest(request)
	_, err := paymentRequests.InsertOne(context.Background(), res)
	return err
}

func GetCompany(companyID string) (res utils.CompanyDTO, err error) {
	filter := bson.M{"_id": companyID}
	companyInstance := company{}
	if err = companies.FindOne(context.Background(), filter).Decode(&companyInstance); err != nil {
		return
	}
	res = companyInstance.convertToCompanyDTO()
	return
}

func GetPaymentRequestsSorted(sortDTO utils.MongoSortDTO) (res []*utils.PaymentRequestDTO, err error) {
	sort := convertToSortOption(sortDTO)
	findOptions := options.Find().SetSort(sort)

	cursor, err := paymentRequests.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var curPayment paymentRequest
		err = cursor.Decode(&curPayment)
		if err != nil {
			return nil, err
		}
		curPaymentDTO := curPayment.convertToPaymentRequestDTO()
		res = append(res, &curPaymentDTO)
	}

	cursor.Close(context.Background())
	return
}

func GetPaymentRequests(filterDTO utils.MongoFilterDTO) (res []*utils.PaymentRequestDTO, err error) {
	filter := convertToFilter(filterDTO)
	cursor, err := paymentRequests.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var curPayment paymentRequest
		err = cursor.Decode(&curPayment)
		if err != nil {
			return nil, err
		}
		curPaymentDTO := curPayment.convertToPaymentRequestDTO()
		res = append(res, &curPaymentDTO)
	}

	cursor.Close(context.Background())
	return
}
