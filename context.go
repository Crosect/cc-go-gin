package ccgin

import (
	"github.com/crosect/cc-go/web/constant"
	"github.com/crosect/cc-go/web/context"
	"github.com/gin-gonic/gin"
)

func InitContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestAttributes := context.GetOrCreateRequestAttributes(c.Request)
		requestAttributes.Mapping = c.FullPath()
		c.Set(constant.ContextReqAttribute, requestAttributes)
		c.Next()
	}
}
