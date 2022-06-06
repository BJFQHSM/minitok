package mysql

type EncryptInfo struct {
	UserId int64  `gorm:"user_id"`
	Salt   string `gorm:"salt"`
}

func (info *EncryptInfo) TableName() string {
	return "encrypt_info"
}

func QueryEncryptInfoByUserId(userId int64) (*EncryptInfo, error) {
	var info EncryptInfo
	err := DB.First(&info, "user_id = ?", userId).Error
	return &info, err
}
