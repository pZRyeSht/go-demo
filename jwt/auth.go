package jwt

import (
	"context"
)

// TokenInfo Token token interface
type TokenInfo interface {
	// GetToken GetXToken 获取访问令牌
	GetToken() string
	// GetTokenType GetXTokenType 获取令牌类型
	GetTokenType() string
	// GetExpiresAt GetXExpiresAt 获取令牌到期时间戳
	GetExpiresAt() int64
	// EncodeToJSON JSON编码
	EncodeToJSON() ([]byte, error)
}

// Author auth interface
type Author interface {
	// Generate TokenGenerate 生成令牌
	Generate(ctx context.Context, userID string) (TokenInfo, error)
	// Destroy TokenDestroy DestroyToken TokenDestroy 销毁令牌
	Destroy(ctx context.Context, xToken string) error
	// ParseUserID ParseUserIDWithToken ParseToken ParseUserID 解析token获得userID
	ParseUserID(ctx context.Context, xToken string) (string, error)
	// Close 释放资源
	Close() error
}

