package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muriiloandrade/memo-app/model"
	"github.com/muriiloandrade/memo-app/model/errors"
)

func (h *Handler) Me(ctx *gin.Context) {
	user, exists := ctx.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context: %v\n", ctx)
		err := errors.InternalError()
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	userId := user.(*model.User).UserId

	user, err := h.UserService.Get(ctx, userId)

	if err != nil {
		log.Printf("Couldn't find user: %v\n%v", user, err)
		e := errors.NotFoundError("user", userId.String())
		ctx.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
