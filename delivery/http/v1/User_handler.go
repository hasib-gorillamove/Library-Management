package v1

import (
	"github.com/golden-infotech/entity"
	"github.com/golden-infotech/lib/logger"
	"github.com/golden-infotech/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
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
	userGroup.GET("/df/:email", h.GetUserByEmail)
	userGroup.PUT("/:id", h.UpdateUser)
	userGroup.DELETE("/:id", h.DeleteUser)
	userGroup.POST("/login", h.Login)
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
func (h *UserHandler) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	res, err := h.UserService.GetUserByEmail(c.Request().Context(), email)
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
		err := c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "invalid request",
		})
		if err != nil {
			return err
		}
	}
	validationError := data.Validate()
	if validationError != nil {
		err := c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Validation Error",
		})
		if err != nil {
			return err
		}
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

func (h *UserHandler) Login(c echo.Context) error {
	login := entity.LoginInfo{}
	err := c.Bind(&login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	user, err := h.UserService.GetUserByEmail(c.Request().Context(), login.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	//checker := lib.HashPassword(login.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Wrong password",
		})
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	token, err := user.GetJwt(expirationTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: "Error signing token",
		})
	}
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   *token,
		Expires: expirationTime,
	})
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "All ok ,Good to go ",
	})
}
