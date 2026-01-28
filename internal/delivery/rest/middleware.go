package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rynr00/go-resto/internal/model/constant"
	"github.com/rynr00/go-resto/internal/tracing"
	"github.com/rynr00/go-resto/internal/usecase/resto"
)

func LoadMiddleware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://go-resto.com"},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))
}

func GetAuthMiddleware(restoUsecase resto.Usecase) *authMiddleware {
	return &authMiddleware{
		restoUsecase: restoUsecase,
	}
}

type authMiddleware struct {
	restoUsecase resto.Usecase
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.CreateSpan(c.Request().Context(), "CheckAuth")
		defer span.End()

		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		userID, err := am.restoUsecase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}
