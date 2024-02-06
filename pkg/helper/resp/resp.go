package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx echo.Context, data interface{}) error {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: 0, Message: "success", Data: data}
	return ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx echo.Context, httpCode, code int, message string, data interface{}) error {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: code, Message: message, Data: data}
	return echo.NewHTTPError(httpCode, resp)
}
