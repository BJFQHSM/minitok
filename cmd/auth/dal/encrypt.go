package dal

type EncryptInfo struct {
	UserId int64  `gorm:"user_id"`
	Salt   string `gorm:"salt"`
}
