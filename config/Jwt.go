package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/shimodii/ecommerce-backend/entities"
)

var PrivateJwtSecret = []byte("mySecret")

func GetSecret() []byte{
    return PrivateJwtSecret
}

func GenerateJwt(user entities.User) (string, error) {
    claims := &jwt.MapClaims{
        "email": user.Email,
        "role": user.Role,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(PrivateJwtSecret)
}
