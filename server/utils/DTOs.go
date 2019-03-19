package utils

type PaymentFromCardDTO struct {
	ID          string `json:"id"`
	CardNumber  string `json:"card_number"`
	CardExpires string `json:"card_expires"`
	CardCVC     string `json:"card_cvc"`
	Amount      int    `json:"amount"`
	Comment     string `json:"comment"`
	Email       string `json:"email"`
	Dangerous   bool   `json:"dangerous"`
}

type PaymentRequestDTO struct {
	ID            string `json:"id,omitempty"`
	INN           string `json:"inn"`
	BIK           string `json:"bik"`
	AccountNumber string `json:"account_number"`
	ForWhat       string `json:"for_what"`
	Amount        int    `json:"amount"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
}

type PaymentViaBankDTO struct {
	INN           string `form:"inn"`
	BIK           string `form:"bik"`
	AccountNumber string `form:"account_number"`
	ForWhat       string `form:"for_what"`
	Amount        int    `form:"amount"`
}

type PatchPaymentFromCardDTO struct {
	Dangerous bool `json:"dangerous" bson:"dangerous"`
}

type CompanyDTO struct {
	Status         string `json:"status"`
	Name           string `json:"name"`
	PhotoPath      string `json:"photo_path"`
	Phone          string `json:"phone"`
	Site           string `json:"site"`
	Email          string `json:"email"`
	Info           string `json:"info"`
	FullInfoPath   string `json:"full_info_path"`
	RequisitesPath string `json:"requisites_path"`
}

type ProductDTO struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
	Price     string `json:"price"`
}

type MongoSortDTO struct {
	Field      string `form:"field"`
	Descending bool   `form:"desc"`
}

type MongoFilterDTO struct {
	Field string `form:"field"`
	Value string `form:"value"`
}
