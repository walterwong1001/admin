package handlers

import "github.com/gin-gonic/gin"

type signInHandler struct{}

func SignInHandler() *signInHandler {
	return &signInHandler{}
}

func (h *signInHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("sign_in", h.SignIn)
}

func (h *signInHandler) SignIn(c *gin.Context) {

}
