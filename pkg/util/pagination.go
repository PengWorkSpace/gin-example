package util

import (
	"github.com/PengWorkSpace/gin-example/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Engine) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
