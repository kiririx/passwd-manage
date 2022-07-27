package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UI(c *gin.Context) {
	// 访问根地址，根据用户的浏览器定位到前端的中间页面，即/ui/#/
	c.Redirect(http.StatusFound, "/ui/#/")
}
