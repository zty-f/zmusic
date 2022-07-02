package dao

type Admin struct {
	ID       uint
	Name     string
	Password string
}

type AdminDao struct {
}

func NewAdminDaoInstance() *AdminDao {
	return &AdminDao{}
}

// QueryAdminByNameAndPassword 通过用户名密码查找用户数
func (u *AdminDao) QueryAdminByNameAndPassword(name, password string) (int64, error) {
	var count int64
	if err := db.Table("admin").Where("name = ? and password = ?", name, password).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
