package db

import (
	"github.com/kadukm/banking_spa/server/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func convertToPaymentFromCard(payment utils.PaymentFromCardDTO) paymentFromCard {
	return paymentFromCard{
		CardNumber:  payment.CardNumber,
		CardExpires: payment.CardExpires,
		CardCVC:     payment.CardCVC,
		Amount:      payment.Amount,
		Comment:     payment.Comment,
		Email:       payment.Email,
		Dangerous:   payment.Dangerous,
	}
}

func (payment paymentFromCard) convertToPaymentFromCardDTO() utils.PaymentFromCardDTO {
	return utils.PaymentFromCardDTO{
		ID:          payment.ID.Hex(),
		CardNumber:  payment.CardNumber,
		CardExpires: payment.CardExpires,
		CardCVC:     payment.CardCVC,
		Amount:      payment.Amount,
		Comment:     payment.Comment,
		Email:       payment.Email,
		Dangerous:   payment.Dangerous,
	}
}

func convertToPaymentRequest(payment utils.PaymentRequestDTO) paymentRequest {
	return paymentRequest{
		INN:           payment.INN,
		BIK:           payment.BIK,
		AccountNumber: payment.AccountNumber,
		ForWhat:       payment.ForWhat,
		Amount:        payment.Amount,
		Phone:         payment.Phone,
		Email:         payment.Email,
	}
}

func (payment paymentRequest) convertToPaymentRequestDTO() utils.PaymentRequestDTO {
	return utils.PaymentRequestDTO{
		ID:            payment.ID.Hex(),
		INN:           payment.INN,
		BIK:           payment.BIK,
		AccountNumber: payment.AccountNumber,
		ForWhat:       payment.ForWhat,
		Amount:        payment.Amount,
		Phone:         payment.Phone,
		Email:         payment.Email,
	}
}

func (company company) convertToCompanyDTO() utils.CompanyDTO {
	return utils.CompanyDTO{
		Status:         company.Status,
		Name:           company.Name,
		PhotoPath:      company.PhotoPath,
		Phone:          company.Phone,
		Site:           company.Site,
		Email:          company.Email,
		Info:           company.Info,
		FullInfoPath:   company.FullInfoPath,
		RequisitesPath: company.RequisitesPath,
	}
}

func (product product) convertToProductDTO() utils.ProductDTO {
	return utils.ProductDTO{
		CompanyID: product.CompanyID,
		Name:      product.Name,
		ImagePath: product.ImagePath,
		Price:     product.Price,
	}
}

func convertToSortOption(sortDTO utils.MongoSortDTO) bson.D {
	if sortDTO.Field == "" {
		return bson.D{}
	}

	filterValue := 1
	if sortDTO.Descending {
		filterValue = -1
	}

	return bson.D{{sortDTO.Field, filterValue}}
}

func convertToFilter(filterDTO utils.MongoFilterDTO) (filter bson.M) {
	if filterDTO.Field == "" {
		return nil
	}
	return bson.M{filterDTO.Field: filterDTO.Value}
}
