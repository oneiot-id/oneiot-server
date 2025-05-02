package helper

import (
	"errors"
	"fmt"
	"oneiot-server/model/entity"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AppClaims defines the custom claims for our JWT
type AppClaims struct {
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

var (
	jwtSecretKey []byte
	jwtExpiresIn time.Duration
)

// LoadJWTConfig loads JWT configuration from environment variables
func LoadJWTConfig() error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return errors.New("JWT_SECRET environment variable not set")
	}
	jwtSecretKey = []byte(secret)

	expMinutesStr := os.Getenv("JWT_EXPIRATION_MINUTES")
	if expMinutesStr == "" {
		expMinutesStr = "1440"
	}
	expMinutes, err := strconv.Atoi(expMinutesStr)
	if err != nil {
		return fmt.Errorf("invalid JWT_EXPIRATION_MINUTES: %w", err)
	}
	jwtExpiresIn = time.Duration(expMinutes) * time.Minute

	return nil
}

// GenerateJWT creates a new JWT string for a given user
func GenerateJWT(user entity.User) (string, time.Time, error) {
	if len(jwtSecretKey) == 0 {
		if err := LoadJWTConfig(); err != nil {
			return "", time.Time{}, fmt.Errorf("failed to load JWT config: %w", err)
		}
	}

	expirationTime := time.Now().Add(jwtExpiresIn)

	claims := &AppClaims{
		UserID:    user.Id,
		UserEmail: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "oneiot-server",       // Identify the issuer
			Subject:   strconv.Itoa(user.Id), // Identify the subject (user ID)
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, expirationTime, nil
}

// ValidateJWT parses and validates a JWT string
func ValidateJWT(tokenString string) (*AppClaims, error) {
	if len(jwtSecretKey) == 0 {
		if err := LoadJWTConfig(); err != nil {
			return nil, fmt.Errorf("failed to load JWT config: %w", err)
		}
	}

	claims := &AppClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("malformed token")
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
