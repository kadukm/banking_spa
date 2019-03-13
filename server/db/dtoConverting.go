package db

import (
	"github.com/kadukm/banking_spa/server/utils"
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
