package controllers

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	service     services.JWTAuthService
	userService services.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	service services.JWTAuthService,
	userService services.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	jwt.logger.Zap.Info("SignIn route called")
	user, _ := jwt.userService.GetOneUser(uint(3))
	token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Zap.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route",
	})
}