package dao

import (
	"time"
)

type Consumer struct {
	ID           uint
	Username     string
	Password     string
	Sex          int8
	PhoneNum     string
	Email        string
	Birth        time.Time
	Introduction string
	Location     string
	Avatar       string
	CreateTime   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ConsumerDao struct {
}

func NewConsumerDaoInstance() *ConsumerDao {
	return &ConsumerDao{}
}

// QueryUserByUserName 根据用户名判断用户是否存在
func (c *ConsumerDao) QueryUserByUserName(username string) (int64, error) {
	var count int64
	if err := db.Table("consumers").Where("username = ?", username).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Add 添加用户
func (c *ConsumerDao) Add(user *Consumer) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
