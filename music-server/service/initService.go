package service

import (
	"zmusic/music-server/dao"
)

const MaxUsernameLen = 32
const MaxPasswordLen = 32

var adminDaoInstance = dao.NewAdminDaoInstance()
