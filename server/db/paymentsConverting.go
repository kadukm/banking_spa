package db

import (
	"errors"

	"github.com/kadukm/banking_spa/server/utils"
)

func convertToPaymentFromCard(payment utils.PaymentFromCardDTO) (res paymentFromCard, err error) {
	if payment.ID != "" {
		return res, errors.New("Given not empty ID of payment from card")
	}
	res = paymentFromCard{
		CardNumber:  payment.CardNumber,
		CardExpires: payment.CardExpires,
		CardCVC:     payment.CardCVC,
		Value:       payment.Value,
		Comment:     payment.Comment,
		Email:       payment.Email,
		Dangerous:   payment.Dangerous,
	}
	return res, nil
}

func convertToPaymentRequest(payment utils.PaymentRequestDTO) (res paymentRequest, err error) {
	if payment.ID != "" {
		return res, errors.New("Given not empty ID of payment request")
	}
	res = paymentRequest{
		Inn:           payment.Inn,
		Bik:           payment.Bik,
		AccountNumber: payment.AccountNumber,
		ForWhat:       payment.ForWhat,
		Value:         payment.Value,
		Phone:         payment.Phone,
		Email:         payment.Email,
	}
	return res, nil
}
