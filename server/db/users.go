package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(login string) (user User, err error) {
	filter := bson.M{"login": login}
	err = users.FindOne(context.TODO(), filter).Decode(&user)
	return
}
