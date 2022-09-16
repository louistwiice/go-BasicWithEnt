package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/domain"
	//"github.com/louistwiice/go/basicwithent/utils"
	"github.com/louistwiice/go/basicwithent/utils/jwt_token"
)


type controller struct {
	service domain.AuthService
}

func NewMiddlewareControllers(svc domain.AuthService) *controller {
	return &controller{
		service: svc,
	}
}

func (cont *controller) JwAuthtMiddleware() gin.HandlerFunc {
	return func (c *gin.Context)  {
		user_id, err := jwttoken.IsTokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. Please Login first", "code": 401, "details": err.Error()})
			return
		}

		user, _, err := cont.service.GetByID(user_id)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. User not found", "code": 401, "details": err.Error()})
			return
		}
		c.Next()
	}
}
