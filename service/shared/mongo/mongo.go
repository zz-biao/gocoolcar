package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ID = "_id"

type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}

// Set return a $set update
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}
