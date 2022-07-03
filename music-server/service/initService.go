package service

import (
	"zmusic/music-server/dao"
)

var adminDaoInstance = dao.NewAdminDaoInstance()
var consumerDaoInstance = dao.NewConsumerDaoInstance()
