package encry

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type EncryKey struct {
	Key        string
	Data       interface{}
	CreateTime time.Time
}

func TestAesCtr(t *testing.T) {
	keys := EncryKey{
		Key:  "mingzhi@ankr.com",
		Data: "!@^#ADJAIOWANDAS",
		CreateTime: time.Now(),
	}

	b, err := json.Marshal(keys)
	assert.NoError(t, err)
	encry, err := AesEncrypt(b)
	assert.NoError(t, err)
	fmt.Println(encry)

	decry, err := AesDecrypt(encry)
	assert.NoError(t, err)
	encry_key := new(EncryKey)
	err = json.Unmarshal(decry, encry_key)
	assert.NoError(t, err)
	fmt.Println(encry_key)
	assert.Equal(t, keys.Key, encry_key.Key)
}
