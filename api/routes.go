package api

import (
	"github.com/gin-gonic/gin"
)

// Routes register api related services into a provided gin.RouterGroup
func Routes(r *gin.RouterGroup) {
	TagRoutes(r)
}
