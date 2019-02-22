package routes

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetTopics(c *gin.Context) {

	c.JSON(200, map[string]interface{}{
		"errno": "0",
	})
}

func PostTopics(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"errno":  "0",
		"errmsg": "post请求提交成功",
	})
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	var user User
	// v := session.Get("user")
	// if v == nil {

	// } else {
 
	// }
	session.Set("user", &user)
	err := session.Save()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, map[string]interface{}{
			"errno":  "-1",
			"errmsg": err.Error(),
		})
	}
	c.JSON(200, gin.H{"user": user})

	// var user User
	// c.BindJSON(&user)
	// session := sessions.Default(c)
	// // 设置一下session,然后返回给浏览器
	// session.Set("user", user)
	// // 保存这个session,登录
	// err := session.Save()
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(200, map[string]interface{}{
	// 		"errno":  "-1",
	// 		"errmsg": err,
	// 	})
	// }
	// c.JSON(200, gin.H{"user": user})
	// if err != nil  {
	// 	c.JSON(200,map[string]interface{}{
	// 		"errno":"-1",
	// 		"errmsg":err,
	// 	})
	// }
	// c.JSON(200, map[string]interface{}{
	// 	"errno":  "0",
	// 	"errmsg": "登录成功",
	// })
}
