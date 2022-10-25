package eth_sign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSign(t *testing.T) {
	message := "hello"
	pri := "49b958f4b453f94e641902db075c1bd9269a4561359daab536652127bc3255f5"
	sign, err := Signature(message, pri)
	assert.Nil(t, err)
	assert.NotNil(t, sign)
	success, err := VerifySignature(message, pri, sign)
	assert.Nil(t, err)
	assert.Equal(t, success, true)
}
