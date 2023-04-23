package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	mgutil "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

type TripRecord struct {
	mgutil.IDField        `bson:"inline"` //inline
	mgutil.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mgutil.NewObjId
	r.UpdatedAt = mgutil.UpdatedAt()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, err
}
