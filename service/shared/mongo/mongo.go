package mgutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedat"
)

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

// Set return a $set update
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
