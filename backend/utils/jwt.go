package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("caelum-secret-key-2024")

// Claims JWT声明
type Claims struct {
	UserID   int64   `json:"user_id"`
	Username string  `json:"username"`
	RoleIDs  []int64 `json:"role_ids"`
	jwt.RegisteredClaims
}

// TokenExpireDuration Token过期时间（小时）
const TokenExpireDuration = 24

// GenerateToken 生成Token
func GenerateToken(userID int64, username string, roleIDs []int64) (string, error) {
	// 设置过期时间
	expireTime := time.Now().Add(TokenExpireDuration * time.Hour)

	// 创建声明
	claims := &Claims{
		UserID:   userID,
		Username: username,
		RoleIDs:  roleIDs,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "caelum",
		},
	}

	// 创建签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken 刷新Token
func RefreshToken(tokenString string) (string, error) {
	// 解析旧token
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// 生成新token
	return GenerateToken(claims.UserID, claims.Username, claims.RoleIDs)
}

// SetJWTSecret 设置JWT密钥
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}