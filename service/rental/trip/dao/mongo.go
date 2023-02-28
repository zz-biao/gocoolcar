package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	mgutil "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("trip"),
		newObjID: primitive.NewObjectID,
	}
}

type TripRecord struct {
	mgutil.IDField        `bson:"inline"`
	mgutil.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {

}
