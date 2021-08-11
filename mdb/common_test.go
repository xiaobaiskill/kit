package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

var (
	testCli *mongo.Client
)

func TestMain(m *testing.M) {
	testCli = NewClient("mongodb://localhost:27017")
	defer testCli.Disconnect(context.Background())
	m.Run()
}
