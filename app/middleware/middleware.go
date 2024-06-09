package middleware

import (
	"goapi/app/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "missing or invalid token")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			tokenStr = authHeader
		}

		claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		c.Set("user", claims)
		return next(c)
	}
}