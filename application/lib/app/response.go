package app

import (
	"captcha/application/lib/retcode"

	"github.com/gin-gonic/gin"
)

// Gin gin返回值结构
type Gin struct {
	C *gin.Context
}

// Response gin 返回值函数
func (g *Gin) Success(data interface{}) {
	g.C.JSON(200, gin.H{
		"code": retcode.SUCCESS,
		"msg":  retcode.GetMsg(retcode.SUCCESS),
		"data": data,
	})
}

// ErrorByCode gin 返回错误
func (g *Gin) ErrorByCode(errCode int) {
	g.C.JSON(200, gin.H{
		"code": errCode,
		"msg":  retcode.GetMsg(errCode),
		"data": "",
	})
	g.C.Abort()
}
