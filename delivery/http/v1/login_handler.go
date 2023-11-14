package v1

import (
	"github.com/golden-infotech/lib/logger"
	"github.com/golden-infotech/service"
)

type LoginHandler struct {
	UserService *service.UserService
	logger      logger.Logger
}

func newLoginHandler(userService *service.UserService, logger logger.Logger) *LoginHandler {
	return &LoginHandler{
		UserService: userService,
		logger:      logger,
	}
}
