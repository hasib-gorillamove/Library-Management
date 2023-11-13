package v1

import (
	"github.com/golden-infotech/entity"
	"github.com/golden-infotech/lib/logger"
	"github.com/golden-infotech/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
	Logger      logger.Logger
}

func NewUserHandler(userService *service.UserService, logger logger.Logger) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Logger:      logger,
	}
}

func (h *UserHandler) MapUserRoutes(userGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	userGroup.POST("", h.CreateUser)
	userGroup.GET("", h.GetAllUser)
	userGroup.GET("/:id", h.GetAUser)
	userGroup.PUT("/:id", h.UpdateUser)
	userGroup.DELETE("/:id", h.DeleteUser)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	data := entity.UserRegistration{}

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Invalid Payload",
			Data:    err,
		})
	}
	validationError := data.Validate()
	if validationError != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Validation Error",
			Data:    err.Error(),
		})
	}
	err = h.UserService.CreateUser(c.Request().Context(), data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "New User Created Successfully",
	})
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	filter := entity.UserFilter{}
	err := c.Bind(&filter)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Internal Server error",
		})
	}
	res, count, err := h.UserService.GetAllUser(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	response := entity.GetAllUserResponses{
		Total: count,
		Page:  filter.Page,
		Users: res,
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "Getting done",
		Data:    response,
	})
}

func (h *UserHandler) GetAUser(c echo.Context) error {
	id := c.Param("id")

	res, err := h.UserService.GetAUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "Successfully get a user",
		Data:    res,
	})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	data := entity.UserRegistration{}
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "invalid request",
		})
	}
	validationError := data.Validate()
	if validationError != nil {
		c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Validation Error",
		})
	}
	err = h.UserService.UpdateUser(c.Request().Context(), data, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "successfully Updated",
	})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := h.UserService.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "successfully deleted User",
	})
}
