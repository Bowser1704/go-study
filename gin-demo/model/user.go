package model

import (
	"github.com/Bowser1704/go-study/gin-demo/pkg/auth"
	"github.com/Bowser1704/go-study/gin-demo/pkg/constvar"
	"fmt"
	validator "gopkg.in/go-playground/validator.v9"
)

// User represents a registered user.
//搞不懂搞个Usermodel干嘛?待解决.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) Tablename() string {
	return "tb_users"	//返回表名字,写个方法.
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error	//Create是gorm.model的方法DB.self是一个*gorm.model在init.go中定义
}

// update user
func (u *UserModel) Update() error {
	return DB.Self.Save(&u).Error
}

func Getuser (id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)	//他用的username但是我用的id
	return u, d.Error
}
func DeleteUser (id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error	//Delete方法应该是靠id(主键)来找的所以只要id就可以了.
}

//查询用户列表,通过名字查询,查询,查询
func ListUser (username string, offest, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprint("username like '%%%s%%'", username)	//转义字符%%%s%%
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(count).Error; err != nil {
		return users, count, err
	}
	if err := DB.Self.Where(where).Offset(offest).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err	//desc mysql中的降序关键字
	}

	return users, count, nil
}

//无尽的方法封装,这里写哪里封装,换来换去....
func (u *UserModel) Compare(pwd string) error {
	err := auth.Compare(u.Password, pwd)
	return err
	//不喜欢他的写法,改了一点
}

//encrypt
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)	//没办法u.Password被定义了
	return err
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
