package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zmusic/music-server/protoc"
)

func AddUser(c *gin.Context) {
	var userAddReq protoc.UserAddReq
	if c.ShouldBind(&userAddReq) == nil {
		log.Printf("userAddReq:{%v}\n", &userAddReq)
	}
	res, err := consumerService.ExistUser(userAddReq.Username)
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
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "用户名已注册，请重新想一个独一无二的吧~",
			Success: false,
			Type:    "warning",
			Data:    nil,
		})
		return
	}
	res, err = consumerService.DoAddUser(&userAddReq)
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
	if !res {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "注册失败，请稍后再试！",
			Success: false,
			Type:    "error",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, protoc.Response{
		Code:    200,
		Message: "注册成功！Welcome~",
		Success: true,
		Type:    "success",
		Data:    nil,
	})
	return
}
