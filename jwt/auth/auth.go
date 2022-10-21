package auth

import (
	"context"
	"errors"
	jwt2 "github.com/EscAlice/go-demo/jwt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	jwtSignKey   = "the-jwt-key"
	jwtTokenType = "Bearer"
	expiresTime  = 7200
)

var (
	ErrInvalidXToken = errors.New("invalid token")
)

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	keyFunc       jwt.Keyfunc
	expired       int64
	tokenType     string
}

// Option 定义参数项
type Option func(*options)

// WithSigningMethod 设定签名方式
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithSigningKey 设定签名key
func WithSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

// WithKeyFunc 设定验证key的回调函数
func WithKeyFunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFunc = keyFunc
	}
}

// WithExpired 设定令牌过期时长(单位秒，默认7200)
func WithExpired(expired int64) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// New 创建认证实例
func New(store Storage, opts ...Option) *JWTAuth {
	o := options{
		signingMethod: jwt.SigningMethodHS256,
		signingKey:    []byte(jwtSignKey),
		expired:       expiresTime,
		tokenType:     jwtTokenType,
		keyFunc: func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidXToken
			}
			return []byte(jwtSignKey), nil
		},
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &JWTAuth{
		opts:  &o,
		store: store,
	}
}

// JWTAuth jwt认证
type JWTAuth struct {
	opts  *options
	store Storage
}

// Generate Token generate 生成令牌
func (a *JWTAuth) Generate(_ context.Context, userID string) (jwt2.TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expired) * time.Second)
	token := jwt.NewWithClaims(a.opts.signingMethod, &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(now),
		Subject:   userID,
	})
	xToken, err := token.SignedString(a.opts.signingKey)
	if err != nil {
		return nil, err
	}
	tokenInfo := &tokenInfo{
		ExpiresAt: expiresAt.Unix(),
		TokenType: a.opts.tokenType,
		Token:     xToken,
	}
	return tokenInfo, nil
}

// parseToken 解析令牌
func (a *JWTAuth) parseToken(xToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(xToken, &jwt.RegisteredClaims{}, a.opts.keyFunc)
	if err != nil || !token.Valid {
		return nil, ErrInvalidXToken
	}

	return token.Claims.(*jwt.RegisteredClaims), nil
}

func (a *JWTAuth) callStore(fn func(storage Storage) error) error {
	if store := a.store; store != nil {
		return fn(store)
	}
	return nil
}

// Destroy Token destroy 销毁令牌
func (a *JWTAuth) Destroy(ctx context.Context, xToken string) error {
	claims, err := a.parseToken(xToken)
	if err != nil {
		return err
	}
	// 则将未过期的令牌放入Storage
	return a.callStore(func(store Storage) error {
		return store.SetKey(ctx, xToken, time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()))
	})
}

// ParseUserID 解析用户ID
func (a *JWTAuth) ParseUserID(ctx context.Context, xToken string) (string, error) {
	if xToken == "" {
		return "", ErrInvalidXToken
	}
	claims, err := a.parseToken(xToken)
	if err != nil {
		return "", err
	}
	err = a.callStore(func(store Storage) error {
		if exists, err := store.CheckKey(ctx, xToken); err != nil {
			return err
		} else if exists {
			return ErrInvalidXToken
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}

// Close 释放资源
func (a *JWTAuth) Close() error {
	return a.callStore(func(store Storage) error {
		return store.Close()
	})
}
