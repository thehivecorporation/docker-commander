package index

import (
    "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    content := gin.H{"Hello": "World"}
    c.JSON(200, content)
}
