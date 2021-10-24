package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

var ormer orm.Ormer
var qs orm.QuerySeter

func init() {
	// 注册模型
	orm.RegisterModel(new(User))
	fmt.Println("注册 User 模型")

	ormer = orm.NewOrm()
	qs = ormer.QueryTable(new(User))
}

type User struct {
	Id       uint64 `orm:"pk,auto"  json:"id"`
	UserName string `orm:"index" json:"userName"`
	Password string `orm:"column(password)" json:"password"`
	Age      int    `json:"age"`
	Addr     string `orm:"null;column(address);size(500)" json:"address"`
	Email    string `orm:"null;size(50)" json:"email"` // 允许为null
}

func (*User) TableName() string {
	return "go_user" // 自定义表名
}

func AddUser(u User) uint64 {
	ormer.Insert(&u)
	return u.Id
}

func GetUser(uid uint64) (u *User, err error) {
	user := User{}
	qs.Filter("id", uid).One(&user)
	//qs.Filter("age__gt",17)
	if user.Id > 0 {
		return &user, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[uint64]*User {
	users := []User{}
	qs.All(&users)

	userMap := make(map[uint64]*User)
	for index, user := range users {
		userMap[user.Id] = &users[index]
	}

	return userMap
}

func GetAllUserList() []User {
	users := []User{}
	qs.All(&users)
	return users
}

func UpdateUser(uid uint64, uu *User) (a *User, err error) {
	user := User{}
	qs.Filter("id", uid).One(&user)
	if user.Id > 0 {
		if uu.Password != "" {
			user.Password = uu.Password
		}
		if uu.Addr != "" {
			user.Addr = uu.Addr
		}
		if uu.Email != "" {
			user.Email = uu.Email
		}
		if uu.Age > 0 {
			user.Age = uu.Age
		}
		ormer.Update(&user)
		return &user, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	user := User{}
	qs.Filter("user_name", username).One(&user)
	if user.Id > 0 {
		if user.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid uint64) {
	user := User{}
	qs.Filter("id", uid).One(&user)
	if user.Id > 0 {
		ormer.Delete(&user)
	}
}
