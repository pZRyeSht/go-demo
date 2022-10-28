package auth

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth(t *testing.T) {
	rd := NewStorage(&RedisConfig{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		Password:     "",
		ReadTimeout:  20,
		WriteTimeout: 20,
	})

	jwtAuth := New(rd)
	defer func(jwtAuth *JWTAuth) {
		err := jwtAuth.Close()
		if err != nil {
			assert.Nil(t, err)
		}
	}(jwtAuth)

	ctx := context.Background()
	userID := "admin"
	token, err := jwtAuth.Generate(ctx, userID)
	assert.Nil(t, err)
	assert.NotNil(t, token)

	id, err := jwtAuth.ParseUserID(ctx, token.GetToken())
	assert.Nil(t, err)
	assert.Equal(t, userID, id)

	err = jwtAuth.Destroy(ctx, token.GetToken())
	assert.Nil(t, err)

	id, err = jwtAuth.ParseUserID(ctx, token.GetToken())
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid token")
	assert.Empty(t, id)
}
