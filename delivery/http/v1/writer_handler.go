package v1

import (
	"fmt"
	"github.com/golden-infotech/entity"
	"github.com/golden-infotech/lib/logger"
	"github.com/golden-infotech/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WriterHandler struct {
	WriterService *service.WriterService
	Logger        logger.Logger
}

func NewWriterHandler(writerService *service.WriterService, logger logger.Logger) *WriterHandler {
	return &WriterHandler{
		WriterService: writerService,
		Logger:        logger,
	}
}

func (h *WriterHandler) MapWriterRoutes(writerGroup *echo.Group, authenticated echo.MiddlewareFunc) {
	writerGroup.POST("", h.Create)
	writerGroup.GET("", h.ListAllWriter)
	writerGroup.GET("/:id", h.GetAWriter)
	writerGroup.PUT("/:id", h.UpdateWriter)
	writerGroup.DELETE("/:id", h.Delete)
}
func (h *WriterHandler) Create(c echo.Context) error {
	data := entity.Writer{}
	err := c.Bind(&data)
	fmt.Println("data...........", data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: "Invalid request payload",
			Data:    err,
		})
	}
	validationErrors := data.Validate()

	if validationErrors != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	err = h.WriterService.Create(c.Request().Context(), data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "New Writer Created Successfully",
	})
}

func (h *WriterHandler) ListAllWriter(c echo.Context) error {
	filter := entity.WriterFilter{}
	err := c.Bind(&filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	res, count, err := h.WriterService.ListAllWriter(c.Request().Context(), filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	response := entity.ListAllWriterResponse{
		Total:  count,
		Page:   filter.Page,
		Writer: res,
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "Getting done",
		Data:    response,
	})
}
func (h *WriterHandler) GetAWriter(c echo.Context) error {
	id := c.Param("id")
	res, err := h.WriterService.GetAWriter(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Successfully get a writer",
		Data:    res,
	})
}

func (h *WriterHandler) UpdateWriter(c echo.Context) error {
	id := c.Param("id")
	data := entity.Writer{}

	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Success: false,
			Message: "Invalid request",
			Data:    err,
		})
	}
	validationError := data.Validate()
	if validationError != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Success: false,
			Message: "validation error",
			Data:    validationError,
		})
	}
	err = h.WriterService.UpdateWriter(c.Request().Context(), data, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.Response{
		Success: true,
		Message: "Successfully Updated",
	})
}

func (h *WriterHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.WriterService.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &entity.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &entity.Response{
		Success: true,
		Message: "Successfully Deleted",
	})
}
