package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"results": "not",
	})
}
