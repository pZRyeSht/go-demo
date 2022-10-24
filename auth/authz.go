package auth

import (
	"context"
	"github.com/EscAlice/go-demo/auth/casbin"
)

type AuthCasbin interface {
	// ParseFromContext parses the user from the context.
	ParseFromContext(ctx context.Context) error
	// GetSub GetSubject returns the subject of the token.
	GetSub() string
	// GetObj GetObject returns the object of the token.
	GetObj() string
	// GetAct GetAction returns the action of the token.
	GetAct() string
	// CheckCasbin ...
	CheckCasbin(ctx context.Context, opts ...casbin.Option) bool
}

type AuthCasbinCreator func() AuthCasbin
