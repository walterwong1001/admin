package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct{}

func Start() {
	r := gin.Default()

	r.Run()
}
