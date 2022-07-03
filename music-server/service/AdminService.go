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

// DoAdminLogin 登录
func (u *AdminService) DoAdminLogin(adminLoginReq *protoc.AdminLoginReq) (bool, error) {
	count, err := adminDaoInstance.QueryAdminByNameAndPassword(adminLoginReq)
	if count == 0 || err != nil {
		log.Printf("[%v]\n", err)
		return false, err
	}
	return true, nil
}
