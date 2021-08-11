package totp

import (
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewKey(t *testing.T) {
	key, err := NewKey("Test", "xiaobaiskill")
	assert.NoError(t, err)
	code, err := totp.GenerateCode(key.Secret(), time.Now())
	assert.NoError(t, err)

	k, err := NewKeyFromUrl(key.URL())
	assert.NoError(t, err)
	assert.True(t, totp.Validate(code, k.Secret()))
}
