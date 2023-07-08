package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(500)" json:"nama_lengkap"`
	Username    string `gorm:"varchar(500)" json:"username"`
	Password    string `gorm:"varchar(500)" json:"password"`
}
