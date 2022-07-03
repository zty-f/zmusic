package service

import "zmusic/music-server/protoc"

type AdminService struct {
}

func NewAdminServiceInstance() *AdminService {
	return &AdminService{}
}

// DoLogin 登录
func (u *AdminService) DoLogin(adminLoginReq *protoc.AdminLoginReq) (bool, error) {
	count, err := adminDaoInstance.QueryAdminByNameAndPassword(adminLoginReq)
	if count == 0 || err != nil {
		return false, err
	}
	return true, nil
}
