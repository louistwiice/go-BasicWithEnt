package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/ent"
)

type RouterBase struct {
	Database		*ent.Client
	OpenApp 		*gin.RouterGroup
}

type Routers struct {
	RouterBase
	RestrictedApp 	*gin.RouterGroup
}