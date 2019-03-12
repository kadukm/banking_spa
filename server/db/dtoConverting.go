package db

import (
	"github.com/kadukm/banking_spa/server/utils"
)

func convertToPaymentFromCard(payment utils.PaymentFromCardDTO) paymentFromCard {
	return paymentFromCard{
		CardNumber:  payment.CardNumber,
		CardExpires: payment.CardExpires,
		CardCVC:     payment.CardCVC,
		Value:       payment.Value,
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
		Value:         payment.Value,
		Phone:         payment.Phone,
		Email:         payment.Email,
	}
}

func (company_ company) convertToCompanyDTO() utils.CompanyDTO {
	return utils.CompanyDTO{
		Status:         company_.Status,
		Name:           company_.Name,
		PhotoPath:      company_.PhotoPath,
		Phone:          company_.Phone,
		Site:           company_.Site,
		Email:          company_.Email,
		Info:           company_.Info,
		FullInfoPath:   company_.FullInfoPath,
		RequisitesPath: company_.RequisitesPath,
	}
}
