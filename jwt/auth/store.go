package auth

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

// Storage 令牌存储接口（通过存储令牌实现主动过期）
type Storage interface {
	// SetKey Store token data and specify the expiration time
	SetKey(ctx context.Context, xToken string, expiration time.Duration) error
	// CheckKey Check whether the token exists
	CheckKey(ctx context.Context, xToken string) (bool, error)
	// Delete ...
	Delete(ctx context.Context, xToken string) (bool, error)
	// Close ...
	Close() error
}

// RedisConfig redis配置参数
type RedisConfig struct {
	Addr         string
	DB           int
	Password     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Store redis存储
type Store struct {
	cli *redis.Client
}

// NewStorage 创建redis存储实例
func NewStorage(cfg *RedisConfig) *Store {
	cli := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		DB:           cfg.DB,
		Password:     cfg.Password,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
	return &Store{
		cli: cli,
	}
}

// SetKey set redis key
func (s *Store) SetKey(_ context.Context, xToken string, expiration time.Duration) error {
	cmd := s.cli.Set(xToken, "1", expiration)
	return cmd.Err()
}

// Delete delete redis key
func (s *Store) Delete(_ context.Context, xToken string) (bool, error) {
	cmd := s.cli.Del(xToken)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// CheckKey Check redis key
func (s *Store) CheckKey(_ context.Context, xToken string) (bool, error) {
	cmd := s.cli.Exists(xToken)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Close ...
func (s *Store) Close() error {
	return s.cli.Close()
}
