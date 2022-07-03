package service

import (
	"log"
	"zmusic/music-server/protoc"
)

type AdminService struct {
}

func NewAdminServiceInstance() *AdminService {
	return &AdminService{}
}

// DoLogin 登录
func (u *AdminService) DoLogin(adminLoginReq *protoc.AdminLoginReq) bool {
	count, err := adminDaoInstance.QueryAdminByNameAndPassword(adminLoginReq)
	if count == 0 || err != nil {
		log.Printf("[%v]\n", err)
		return false
	}
	return true
}
