package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct{}

// Config will hold services that will eventually be injected
// into this handler layer on handler initialization
type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	handler := &Handler{}

	group := c.R.Group(os.Getenv("ACCOUNT_API_URL"))

	group.GET("/me", handler.Me)
	group.POST("/signup", handler.SignUp)
	group.POST("/signin", handler.SignIn)
	group.POST("/signout", handler.SignOut)
	group.POST("/tokens", handler.Tokens)
	group.POST("/image", handler.Image)
	group.DELETE("/image", handler.DeleteImage)
	group.PUT("/details", handler.Details)
}

func (h *Handler) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

func (h *Handler) SignOut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

func (h *Handler) Tokens(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

func (h *Handler) Image(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}

func (h *Handler) DeleteImage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

func (h *Handler) Details(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
