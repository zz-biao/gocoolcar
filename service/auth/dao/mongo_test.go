package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://root:123456@localhost:27017/?authSource=admin"))
	if err != nil {
		t.Fatalf("cannot connnect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resolve acount id for 123: %v", err)
	} else {
		want := "123444" //todo
		if id != want {
			t.Errorf("resolve acount id: want: %q, got: %q", want, id)
		}
	}
}
