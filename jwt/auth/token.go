package auth

import "encoding/json"

// tokenInfo 令牌信息
type tokenInfo struct {
	Token     string `json:"token"`      // 访问令牌
	TokenType string `json:"token_type"` // 令牌类型
	ExpiresAt int64  `json:"expires_at"` // 令牌到期时间
}

func (t *tokenInfo) GetToken() string {
	return t.Token
}

func (t *tokenInfo) GetTokenType() string {
	return t.TokenType
}

func (t *tokenInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t *tokenInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}
