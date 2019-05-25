package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name		int
	Avatar		string
	StNum		string		`gorm:"unique"`
	Tel			string
	QQ			string
	Wechat		string

}
