package rdb

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	t1, err := time.Parse(timeLayout, "0001-01-01 00:00:00")
	assert.NoError(t, err)
	t.Log(t1.Unix())
}

func TestJson(t *testing.T) {
	type data struct {
		CreateTime Time `json:"create_time"`
	}
	d := &data{FromTime(time.Now())}
	b, err := json.Marshal(d)
	t.Log(string(b))
	assert.NoError(t, err)
	d1 := &data{}
	err = json.Unmarshal(b, d1)
	assert.NoError(t, err)
	fmt.Printf("%v", d1)
}
