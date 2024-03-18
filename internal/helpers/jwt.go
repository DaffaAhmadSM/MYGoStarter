package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Siswa struct {
	Id     int
	Name   string
	RoleId int
	jwt.RegisteredClaims
}

func IssueToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		cookie, err := ec.Cookie("token")
		if err != nil {
			return ec.Redirect(http.StatusUnauthorized, "/")
		}
		token, _ := jwt.ParseWithClaims(cookie.Value, &Siswa{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if !token.Valid {
			return ec.Redirect(http.StatusUnauthorized, "/")
		}
		return next(ec)
	}
}
