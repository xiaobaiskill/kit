package tokenverfiy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswd(t *testing.T) {
	p := NewTokenVerfiy("aaa")
	secret := "jmz"
	token := p.GenerateToken(secret)
	assert.True(t, p.Verfiy(token, secret))
}
