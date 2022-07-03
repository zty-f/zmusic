package consumer

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
