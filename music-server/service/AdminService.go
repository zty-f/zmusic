package service

type AdminService struct {
}

func NewAdminServiceInstance() *AdminService {
	return &AdminService{}
}

// DoLogin 登录
func (u *AdminService) DoLogin(name, password string) (bool, error) {
	count, err := adminDaoInstance.QueryAdminByNameAndPassword(name, password)
	if count == 0 || err != nil {
		return false, err
	}
	return true, nil
}
