package user

import (
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type Controller struct{}

// Index action: GET /users
func (pc Controller) HelloIndex(c *gin.Context) {
	c.JSON(200, "Helloworld")
}
