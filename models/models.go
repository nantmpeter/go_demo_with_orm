package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/personal?charset=utf8")
	orm.RegisterModel(new(User), new(Profile), new(Post))
	createTable()
}

type User struct {
	Id       int
	Username string
	Password string
	Profile  *Profile `orm:"rel(one)"`
}

type Profile struct {
	Id      int
	Gender  string
	Age     int
	Address string
	Email   string
	User    *User   `orm:"reverse(one)"`
}

type Post struct {
	Id int
	Title string
	Create_at int
	Update_at int
	Model int //类型
	content string
	User *User `orm:"rel(fk)"`
}

func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}

