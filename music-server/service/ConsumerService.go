package service

import (
	"log"
	"time"
	"zmusic/music-server/dao"
	"zmusic/music-server/protoc"
)

type ConsumerService struct {
}

func NewConsumerServiceInstance() *ConsumerService {
	return &ConsumerService{}
}

// ExistUser 判断用户名是否存在
func (u *ConsumerService) ExistUser(username string) (bool, error) {
	count, err := consumerDaoInstance.QueryUserByUserName(username)
	if count == 0 || err != nil {
		log.Printf("[%v]\n", err)
		return false, err
	}
	return true, nil
}

// DoAddUser 添加用户
func (u *ConsumerService) DoAddUser(userAddReq *protoc.UserAddReq) (bool, error) {
	birthTime, _ := time.Parse("2006-01-02", userAddReq.Birth)
	user := &dao.Consumer{
		Username:     userAddReq.Username,
		Password:     userAddReq.Password,
		Sex:          int8(userAddReq.Sex),
		PhoneNum:     userAddReq.PhoneNum,
		Email:        userAddReq.Email,
		Birth:        birthTime,
		Introduction: userAddReq.Introduction,
		Location:     userAddReq.Location,
		Avatar:       "/",
	}
	err := consumerDaoInstance.Add(user)
	if err != nil {
		log.Printf("[%v]\n", err)
		return false, err
	}
	return true, nil
}

// DoUserLogin 用户登录
func (u *ConsumerService) DoUserLogin(userLoginReq *protoc.UserLoginReq) (bool, error) {
	count, err := consumerDaoInstance.QueryUserByNameAndPassword(userLoginReq)
	if count == 0 || err != nil {
		log.Printf("[%v]\n", err)
		return false, err
	}
	return true, nil
}

// GetUserByUserName 根据用户名查找用户
func (u *ConsumerService) GetUserByUserName(username string) (*dao.Consumer, error) {
	user, err := consumerDaoInstance.QueryUserByName(username)
	if err != nil {
		log.Printf("[%v]\n", err)
		return nil, err
	}
	return user, nil
}
