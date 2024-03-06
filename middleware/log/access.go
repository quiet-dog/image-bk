package log

import (
	"fmt"
	"image-bk/global"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := color.New(color.BgGreen).SprintFunc()(" " + c.Request.URL.Path + " ")
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)

		var status string
		if c.Writer.Status() != 200 {
			// status = color.RedString("%d", c.Writer.Status())
			// 设置背景色
			status = color.New(color.BgRed).SprintFunc()(" " + fmt.Sprintf("%d", c.Writer.Status()) + " ")
		} else {
			// status = color.GreenString("%d", c.Writer.Status())
			status = color.New(color.BgGreen).SprintFunc()(" " + fmt.Sprintf("%d", c.Writer.Status()) + " ")
		}

		var method string
		if c.Request.Method == "GET" {
			method = color.New(color.BgBlue).SprintFunc()(" " + c.Request.Method + " ")
		} else {
			method = color.New(color.BgYellow).SprintFunc()(" " + c.Request.Method + " ")
		}

		clientIP := color.New(color.BgCyan).SprintFunc()(" " + c.ClientIP() + " ")

		// 输出颜色
		global.LOG.Info(
			path,
			"  "+status,
			"  "+method+"  ",
			//zap.String("path", path),
			query+"  ",
			clientIP+"  ",
			c.Request.UserAgent()+"  ",
			c.Errors.ByType(gin.ErrorTypePrivate).String()+"  ",
			cost,
		)
	}
}
