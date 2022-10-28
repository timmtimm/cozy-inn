package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whitelist []string = make([]string, 5)

type JwtCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

func (jwtConf *ConfigJWT) GenerateToken(email string, role string) string {
	claims := JwtCustomClaims{
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtConf.SecretJWT))

	whitelist = append(whitelist, token)

	return token
}

func (jwtConf *ConfigJWT) GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)

	_, err := jwtConf.CheckToken(user.Raw)

	if err != nil {
		return nil
	}

	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

func (jwtConf *ConfigJWT) CheckToken(token string) (*JwtCustomClaims, error) {
	for _, tkn := range whitelist {
		if tkn == token {
			token, err := jwt.ParseWithClaims(
				token,
				&JwtCustomClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtConf.SecretJWT), nil
				},
			)
			if err != nil {
				return nil, err
			}

			claims, ok := token.Claims.(*JwtCustomClaims)
			if !ok {
				return nil, errors.New("couldn't parse claims")
			}

			if claims.ExpiresAt < time.Now().Local().Unix() {
				return nil, errors.New("token expired")
			}

			return claims, nil
		}
	}

	return nil, errors.New("token not found")
}

func Logout(token string) bool {
	for idx, tkn := range whitelist {
		if tkn == token {
			whitelist = append(whitelist[:idx], whitelist[idx+1:]...)
		}
	}

	return true
}
