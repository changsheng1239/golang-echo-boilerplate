package handler

import (
	"echo-boilerplate/internal/service"
	"echo-boilerplate/pkg/helper/resp"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUserByID(ctx echo.Context) error
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(h *Handler, us service.UserService) UserHandler {
	return &userHandler{h, us}
}

func (uh *userHandler) GetUserByID(ctx echo.Context) error {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, "invalid id", nil)
	}

	user, err := uh.userService.GetUserByID(id)
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, user)
}
