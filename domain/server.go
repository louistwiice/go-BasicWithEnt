package domain

import "github.com/gin-gonic/gin"

type ServerMiddleware interface {
	JwAuthtMiddleware() gin.HandlerFunc
}