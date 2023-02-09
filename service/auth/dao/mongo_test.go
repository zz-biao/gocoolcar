package dao

import (
	"context"
	mongotesting "coolcar/shared/testing"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	//start container
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
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

	//remove container
}

func TestMain(m *testing.M) {

	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
