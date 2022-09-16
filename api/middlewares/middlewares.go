package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/domain"
	"github.com/louistwiice/go/basicwithent/utils/jwt_token"
)


type controller struct {
	middlewareController domain.ServerMiddleware
}

func NewMiddlewareControllers() *controller {
	return &controller{}
}

func (cont *controller) JwAuthtMiddleware() gin.HandlerFunc {
	return func (c *gin.Context)  {
		err := jwttoken.IsTokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. Please Login first", "code": "401", "details": err.Error()})
			return
		}
		c.Next()
	}
}
