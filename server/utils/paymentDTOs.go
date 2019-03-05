package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentFromCardDTO struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CardNumber  string             `json:"card_number" bson:"card_number"`
	CardExpires string             `json:"card_expires" bson:"card_expires"`
	CardCVC     int                `json:"card_cvc" bson:"card_cvc"`
	Value       int                `json:"value"`
	Comment     string             `json:"comment"`
	Email       string             `json:"email"`
	Dangerous   bool               `json:"dangerous"`
}

type PaymentRequestDTO struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Inn           int                `json:"inn"`
	Bik           int                `json:"bik"`
	AccountNumber string             `json:"account_number"`
	ForWhat       string             `json:"for_what"`
	Value         int                `json:"value"`
	Phone         string             `json:"phone"`
	Email         string             `json:"email"`
}

type PaymentViaBankDTO struct {
	Inn           int    `json:"inn"`
	Bik           int    `json:"bik"`
	AccountNumber string `json:"account_number"`
	ForWhat       string `json:"for_what"`
	Value         int    `json:"value"`
}

type PatchPaymentFromCardDTO struct {
	Dangerous bool `json:"dangerous" bson:"dangerous"`
}
