package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
)

type contextKey string

const (
	tokenPayload contextKey = "payload"
)

type currency string

const (
	BTC currency = "btcusdt@trade"
	ETH currency = "ethusdt@trade"
	SOL currency = "solusdt@trade"
)

type SignUpUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7"`
}

type SignUpUserResponse struct {
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	UserID   int64  `json:"user_id" validate:"required,number,min=1"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7"`
}

type LoginUserResponse struct {
	AccessToken          string             `json:"access_token"`
	AccessTokenExpiresAt time.Time          `json:"access_token_expires_at"`
	User                 SignUpUserResponse `json:"user"`
}

var (
	ErrTokenExpired        = errors.New("token has expired")
	ErrInvalidToken        = errors.New("token is invalid")
	ErrNoAuthHeader        = errors.New("no authorization header")
	ErrInvalidAuthHeader   = errors.New("invalid authorization header")
	ErrUnsupportedAuthType = errors.New("unsupported authorization type")
	ErrInvalidKeySize      = fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	ErrBadRequest          = errors.New("bad request")
	ErrNotAuthorized       = errors.New("not authorized")
	ErrSubscriptionFailed  = errors.New("subscription failed")
)
