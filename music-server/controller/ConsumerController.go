package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zmusic/music-server/protoc"
)

type UserLoginResp struct {
	protoc.Response
	Data *protoc.UserVo `json:"data,omitempty"`
}

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
		})
		return
	}
	if res {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "用户名已注册，请重新想一个独一无二的吧~",
			Success: false,
			Type:    "warning",
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
		})
		return
	}
	if !res {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "注册失败，请稍后再试！",
			Success: false,
			Type:    "error",
		})
		return
	}
	c.JSON(http.StatusOK, protoc.Response{
		Code:    200,
		Message: "注册成功！Welcome~",
		Success: true,
		Type:    "success",
	})
	return
}

func UserLogin(c *gin.Context) {
	var userLoginReq protoc.UserLoginReq
	if c.ShouldBind(&userLoginReq) == nil {
		log.Printf("userLoginReq:{%v}\n", &userLoginReq)
	}
	if userLoginReq.Username == "" || userLoginReq.Password == "" {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    200,
			Message: "请输入用户名和密码！",
			Success: false,
			Type:    "error",
		})
		return
	}
	fmt.Printf("登录的用户为：[%v]", &userLoginReq)
	//调用Service层
	res, err := consumerService.DoUserLogin(&userLoginReq)
	if err != nil {
		c.JSON(http.StatusOK, protoc.Response{
			Code:    500,
			Message: err.Error(),
			Success: false,
			Type:    "error",
		})
		return
	}
	if res {
		c.Set("username", userLoginReq.Username)
		user, err1 := consumerService.GetUserByUserName(userLoginReq.Username)
		userVo := &protoc.UserVo{
			Id:       int32(user.ID),
			Username: user.Username,
			PhoneNum: user.Password,
			Email:    user.Email,
			Avatar:   user.Avatar,
		}
		if err1 != nil {
			c.JSON(http.StatusOK, protoc.Response{
				Code:    500,
				Message: err1.Error(),
				Success: false,
				Type:    "error",
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResp{
			Response: protoc.Response{
				Code:    200,
				Message: "登录成功！",
				Success: true,
				Type:    "success",
			},
			Data: userVo,
		})
		return
	}
	c.JSON(http.StatusOK, protoc.Response{
		Code:    200,
		Message: "用户名或者密码错误",
		Success: false,
		Type:    "error",
	})
	return
}
