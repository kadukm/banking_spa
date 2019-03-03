package utils

type PaymentFromCardDTO struct {
	ID          uint64 `json:"id"`
	CardNumber  string `json:"card_number"`
	CardExpires string `json:"card_expires"`
	CardCVC     int    `json:"card_cvc"`
	Value       int    `json:"value"`
	Comment     string `json:"comment"`
	Email       string `json:"email"`
	Dangerous   bool   `json:"dangerous"`
}

type PaymentRequestDTO struct {
	ID            uint64 `json:"id"`
	Inn           int    `json:"inn"`
	Bik           int    `json:"bik"`
	AccountNumber string `json:"account_number"`
	ForWhat       string `json:"for_what"`
	Value         int    `json:"value"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
}

type PaymentViaBankDTO struct {
	Inn           int    `json:"inn"`
	Bik           int    `json:"bik"`
	AccountNumber string `json:"account_number"`
	ForWhat       string `json:"for_what"`
	Value         int    `json:"value"`
}

type PaymentFromCardPatch struct {
	Dangerous bool `json:"dangerous"`
}
