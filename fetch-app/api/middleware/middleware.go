package middleware

import (
	"net/http"
	"strings"

	"monorepo/common"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS256
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.NewMissingHeaderResponse("Authorization"))
			}
			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.NewInternalServerError("invalid token"))
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, nil)

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSigningMethod {
				return c.JSON(http.StatusForbidden, common.NewInternalServerError("Signing method invalid"))
			}

			return next(c)
		}
	}
}

func JWTMiddlewareAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, common.NewMissingHeaderResponse("Authorization"))
			}
			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, common.NewInternalServerError("invalid token"))
			}

			claim := jwt.MapClaims{}
			token, _ := jwt.ParseWithClaims(signature[1], claim, nil)

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwtSigningMethod {
				return c.JSON(http.StatusForbidden, common.NewInternalServerError("Signing method invalid"))
			}

			data := claim["user"]
			role, _ := data.(map[string]interface{})
			admin := role["role"].(string)

			if admin != "admin" {
				return c.JSON(http.StatusForbidden, common.NewInternalServerError("Failed on verify role, role must admin"))
			}

			return next(c)
		}
	}
}
