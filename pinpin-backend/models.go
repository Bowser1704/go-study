package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name		string
	Avatar		string
	StNum		string		`gorm:"unique"`
	Tel			string
	QQ			string
	Wechat		string

}

type Order struct {
	gorm.Model
	Time		time.Time
	Content 	string
	PostID 		uint
	Users		[]* User	`gorm:"many2many:user_orders;"`
}

func openDB(username, password, addr, dbname string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		dbname,
		true,
		"Asia%2FShanghai")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Print(err)
	}

	return db
}

func main() {
	db := openDB("root", "@Bowser1704", "127.0.0.1:3306", "pinpin")
	db.AutoMigrate(&User{})
	if !db.HasTable(&User{}){
		db.CreateTable(&User{})
	}

	//user := User{Name:"Bowser",
	//	Avatar:"https://avatars2.githubusercontent.com/u/43539191?s=460&v=4",
	//	StNum:"2018212576",
	//	QQ:"896379346",
	//	Tel:"111111111",
	//	Wechat:"wwwwww",
	//}
	//db.Create(&user)
	db.Close()
}