package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rynr00/go-resto/internal/model"
	"github.com/rynr00/go-resto/internal/model/constant"
	"github.com/rynr00/go-resto/internal/tracing"
	"github.com/sirupsen/logrus"
)

func (h *handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()

	var request model.OrderMenuRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("get error: %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

	orderData, err := h.restoUsecase.Order(ctx, request)
	if err != nil {
		fmt.Printf("get error: %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetOrderInfo")
	defer span.End()

	orderID := c.Param("orderID")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUsecase.GetOrderInfo(ctx, model.GetOrderInfoRequest{
		UserID:  userID,
		OrderID: orderID,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("[delivery][rest][order_handler][GetOrderInfo] unable to get order data")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
