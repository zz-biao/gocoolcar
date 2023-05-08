package dao

import (
	"context"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDFiled = "open_id"

type Mongo struct {
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("account"),
		newObjID: primitive.NewObjectID,
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, openID string) (id.AccountID, error) {

	insertID := m.newObjID
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDFiled: openID,
	}, mgutil.SetOnInsert(bson.M{
		mgutil.IDFieldName: insertID,
		openIDFiled:        openID,
	}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", err)
	}

	var row mgutil.IDField
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("res Decode: %v", err)
	}
	return objid.ToAccountID(row.ID), nil
}
