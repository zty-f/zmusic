package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zmusic/music-server/protoc"
)

func AdminLogin(c *gin.Context) {
	var adminLoginReq protoc.AdminLoginReq
	if c.ShouldBind(&adminLoginReq) == nil {
		log.Printf("adminLoginReq:{%v}\n", &adminLoginReq)
	}
	if adminLoginReq.Name == "" || adminLoginReq.Password == "" {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "请输入用户名和密码！",
			Success: false,
			Type:    "error",
			Data:    nil,
		})
		return
	}
	fmt.Println("登录的管理员为：" + adminLoginReq.Name + ":" + adminLoginReq.Password)
	//调用Service层
	res, err := adminService.DoLogin(&adminLoginReq)
	if err != nil {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    500,
			Message: err.Error(),
			Success: false,
			Type:    "error",
			Data:    nil,
		})
		return
	}
	if res {
		c.Set("name", adminLoginReq.Name)
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "登录成功！",
			Success: true,
			Type:    "success",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, protoc.Response{
		Code:    200,
		Message: "用户名或者密码错误",
		Success: false,
		Type:    "error",
		Data:    nil,
	})
	return
}
