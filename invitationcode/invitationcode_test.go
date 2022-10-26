package invitationcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeInviteCode(t *testing.T) {
	var inviteID uint64 = 1
	code := EncodeInviteCode(inviteID)
	assert.NotEqual(t, "", code)
	actInviteID := DecodeInviteCode(code)
	assert.Equal(t, inviteID, actInviteID)
}
