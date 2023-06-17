package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	// 设置静态文件目录（可选）
	router.Static("/static", "./static")
	// 设置HTML模板文件目录（可选）
	router.LoadHTMLGlob("view/*")

	// userApi := router.Group("user/api")

	// 獲得SQL injection 網頁
	router.GET("/SQLinjection", func(c *gin.Context) {
		c.HTML(200, "SQLinjection.html", gin.H{
			"message" : "done",
		})
	})

	// 獲得SQL injection 網頁
	router.POST("/SQLinjection", func(c *gin.Context) {
		// username := c.PostForm("name")
		// password := c.PostForm("password")
		// sql := "SELECT * FROM user WHERE username='"+username+"' AND password='"+password+"'"
		// rows, err := database.DB.Exec(sql)
		
		c.HTML(200, "SQLinjection.html", gin.H{
			"message" : "done",
		})
	})
}


