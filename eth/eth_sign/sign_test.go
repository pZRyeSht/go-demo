package eth_sign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSign(t *testing.T) {
	signature := "0xecf1dae4684ad9448350626f10d6086ca08fd71d77744dcc8851985adc191f0b42d444bed8c4bdf9bdc7a42bcaa26039f39b37e43334bd3cda2fceaf9a905e9a1b"
	message := "Decentraland Login\nEphemeral address: 0xD52d1c7410D79b7FA1f210EEbAdff53312325Fbf\nExpiration: 2022-11-03T08:35:55.176Z"
	privateKey := "afcd4f7678eee997c9194c99237b49b0910a6c936a6233437313a05426984f36"
	address := "0x66d3b64db8d7461f40b1a28bc2d0b494f1bff017"
	sign, err := PersonalSignature(message, privateKey)
	assert.Nil(t, err)
	assert.Equal(t, sign, signature)
	success, err := VerifySignature(message, address, signature)
	assert.Nil(t, err)
	assert.Equal(t, success, true)
}
