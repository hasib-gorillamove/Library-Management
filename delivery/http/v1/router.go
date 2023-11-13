package v1

import (
	"github.com/golden-infotech/config"
	"github.com/golden-infotech/lib/logger"
	"github.com/golden-infotech/repository"
	"github.com/golden-infotech/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

// Setup all routers
func SetupRouters(c *echo.Echo, conf *config.Config, db *bun.DB, jwtConfig middleware.JWTConfig, logger logger.Logger) {

	booksRepository := repository.NewBooksRepository(db)
	booksService := service.NewBooksService(booksRepository)
	booksHandler := NewBooksHandler(booksService, logger)

	writerRepository := repository.NewWriterRepository(db)
	writerService := service.NewWriterService(writerRepository)
	writerHandler := NewWriterHandler(writerService, logger)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService, logger)

	authenticated := middleware.JWTWithConfig(jwtConfig)

	v1 := c.Group("/api/v1")

	booksGroup := v1.Group("/books")
	writerGroup := v1.Group("/writer")
	userGroup := v1.Group("/user_registration")

	booksHandler.MapBooksRoutes(booksGroup, authenticated)
	writerHandler.MapWriterRoutes(writerGroup, authenticated)
	userHandler.MapUserRoutes(userGroup, authenticated)

}
