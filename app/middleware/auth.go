package middleware

import (
	"cozy-inn/util"
	"errors"
	"net/http"
	"strings"
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

type RoleMiddleware struct {
	Role []string
	Func echo.HandlerFunc
}

func (rm RoleMiddleware) CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", -1)

		found := false
		index := 0

		for idx, tkn := range whitelist {
			if tkn == token {
				index = idx
				found = true
			}
		}

		if found {
			t, err := jwt.ParseWithClaims(
				token,
				&JwtCustomClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(util.GetConfig("JWT_SECRET_KEY")), nil
				},
			)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid token",
				})
			}

			claims, ok := t.Claims.(*JwtCustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("Unauthorized"))
			}

			if claims.ExpiresAt < time.Now().Local().Unix() {
				whitelist = append(whitelist[:index], whitelist[index+1:]...)
				return echo.NewHTTPError(http.StatusUnauthorized, errors.New("token expired"))
			}

			for _, role := range rm.Role {
				if claims.Role == role {
					return rm.Func(c)
				} else {
					return c.JSON(http.StatusUnauthorized, map[string]string{
						"message": "forbidden",
					})
				}
			}
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "unauthorized",
		})
	}
}

func Logout(token string) bool {
	for idx, tkn := range whitelist {
		if tkn == token {
			whitelist = append(whitelist[:idx], whitelist[idx+1:]...)
		}
	}

	return true
}
