package route

import (
	"net/http"

	"github.com/zu1k/she/processor"

	"github.com/gin-gonic/gin"
)

func search(c *gin.Context) {
	key := c.Query("key")
	mode := c.DefaultQuery("mode", "1")
	switch mode {
	case "0": // result结构
		result := processor.Search(key)
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	case "1": // 默认 json输出
		result := processor.Search2Json(key)
		c.String(http.StatusOK, result)
	}
}
