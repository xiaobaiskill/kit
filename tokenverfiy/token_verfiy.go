package tokenverfiy

import (
	"crypto/md5"
	"encoding/hex"
)

type TokenVerfiy struct {
	Key string
}

func NewTokenVerfiy(key string) *TokenVerfiy {
	return &TokenVerfiy{Key: key}
}

func (t *TokenVerfiy) Verfiy(token string, secret string) bool {
	ctx := md5.New()
	ctx.Write([]byte(secret))
	return hex.EncodeToString(ctx.Sum([]byte(t.Key))) == token
}

func (t *TokenVerfiy) GenerateToken(secret string) string {
	ctx := md5.New()
	ctx.Write([]byte(secret))
	return hex.EncodeToString(ctx.Sum([]byte(t.Key)))
}
