package password

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerfiy(t *testing.T) {
	Get()
	passwd = &passwdConfig{
		Key:   "ankr",
		Token: "616e6b72ac5c2983f2860c5b340151053598bf87",
	}
	assert.True(t, Verfiy("jmz"))

	assert.False(t, Verfiy("jmz1"))
}

func TestGeneric(t *testing.T) {
	ctx := md5.New()
	ctx.Write([]byte("passwd"))
	t.Log(hex.EncodeToString(ctx.Sum([]byte("key"))))
}
