package route

import (
	"net/http"

	"github.com/zu1k/she/persistence"

	"github.com/gin-gonic/gin"
)

func getSource(c *gin.Context) {
	sources, err := persistence.FetchAllSource()
	if err != nil || len(sources) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"sources": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"sources": sources,
		})
	}
}
