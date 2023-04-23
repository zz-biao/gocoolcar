package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
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

	trip, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
	})
	if err != nil {
		return
	}
	t.Errorf("", trip) //log logf 可能需要程序出错才可以看到
	//remove container
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
