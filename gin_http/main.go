package main 

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
}

func main() {
	// 1. 创建路由
	router := gin.Default()

	// gin.Context 封装了response和request
	router.LoadHTMLGlob("html/*") 
	/* 加载html模板的函数
	1.router.LoadHTMLGlob("templates/*") 
	这个函数可以使用通配符，即*、? 和 []
	2. router.LoadHTMLFiles
	这个函数不可以使用通配符，必须指定清楚要加载的文件名
	*/
	router.GET("/", sayHello)
	// 2. 绑定路由规则，执行函数	
	router.POST("/login", getUser)
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello World!")
	// })

	// 3. 监听端口，默认8080
	// Run("里面不指定端口号默认为8080") 
	router.Run(":8000")
}

func sayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getUser(c *gin.Context) {
	var user User
	// fmt.Println(c.PostForm("username"))
	// fmt.Println(c.PostForm("password"))
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	fmt.Println("user = ", user)
	c.String(http.StatusOK, "Login Successful")
}