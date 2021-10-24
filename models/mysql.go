package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	username, _ := beego.AppConfig.String("username")
	password, _ := beego.AppConfig.String("password")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	database, _ := beego.AppConfig.String("database")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, database)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
	fmt.Println("初始化数据库")
}
