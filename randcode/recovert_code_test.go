package randcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	codes := GenerateRecoveryCode(5, 16)
	assert.Equal(t, len(codes), 16)
	fmt.Println(codes)
}
