//+build integration

package mdb

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaobaiskill/kit/mdb/test"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

// set MONGO_URL env
func TestNewMongoRepository(t *testing.T) {
	r := NewRepository(testCli, "test", "test")
	type foo struct {
		Bar string `bson:"bar"`
	}
	test.DoTestInTx(testCli, func(ctx mongo.SessionContext) {
		err := r.AddOne(ctx, &foo{Bar: "bar"})
		assert.NoError(t, err)
	})
}
