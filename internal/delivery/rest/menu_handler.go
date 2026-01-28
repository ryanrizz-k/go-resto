package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rynr00/go-resto/internal/tracing"
)

func (h *handler) GetMenuList(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetMenuList")
	defer span.End()

	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenuList(ctx, menuType)
	if err != nil {
		fmt.Printf("get error: %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
