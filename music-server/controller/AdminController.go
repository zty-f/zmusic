package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zmusic/music-server/protoc"
)

func AdminLogin(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "请输入用户名和密码！",
			Success: false,
			Type:    "error",
			Data:    nil,
		})
		return
	}
	fmt.Println("登录的管理员为：" + name + ":" + password)
	//调用Service层
	res, err := adminService.DoLogin(name, password)
	if res && err == nil {
		c.Set("name", name)
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "登录成功！",
			Success: true,
			Type:    "success",
			Data:    nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "用户名或者密码错误",
			Success: false,
			Type:    "error",
			Data:    nil,
		})
	}
	return
}
