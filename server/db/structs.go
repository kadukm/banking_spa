package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type paymentFromCard struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CardNumber  string             `bson:"card_number"`
	CardExpires string             `bson:"card_expires"`
	CardCVC     int                `bson:"card_cvc"`
	Value       int                `bson:"value"`
	Comment     string             `bson:"comment"`
	Email       string             `bson:"email"`
	Dangerous   bool               `bson:"dangerous"`
}

type paymentRequest struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	INN           string             `bson:"inn"`
	BIK           string             `bson:"bik"`
	AccountNumber string             `bson:"account_number"`
	ForWhat       string             `bson:"for_what"`
	Value         int                `bson:"value"`
	Phone         string             `bson:"phone"`
	Email         string             `bson:"email"`
}

type company struct {
	ID             string `bson:"_id"`
	Status         string `bson:"status"`
	Name           string `bson:"name"`
	PhotoPath      string `bson:"photo_path"`
	Phone          string `bson:"phone"`
	Site           string `bson:"site"`
	Email          string `bson:"email"`
	FullInfoPath   string `bson:"full_info_path"`
	RequisitesPath string `bson:"requisites_path"`
}
