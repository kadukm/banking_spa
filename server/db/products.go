package db

import (
	"context"

	"github.com/kadukm/banking_spa/server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(companyID string, maxCount int64) (res []*utils.ProductDTO, err error) {
	filter := bson.M{"company_id": companyID}
	findOptions := options.Find().SetLimit(maxCount)

	cursor, err := products.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var curProduct product
		err = cursor.Decode(&curProduct)
		if err != nil {
			return nil, err
		}
		curProductDTO := curProduct.convertToProductDTO()
		res = append(res, &curProductDTO)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.Background())
	return
}
