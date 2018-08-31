package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Users struct{
	Id 		int		`json:"id"`
	Name 	string	`json:"name"`
	Mobile  string	`json:"mobile"`
}

func init(){
	orm.RegisterDataBase(
		"default",
		"mysql",
		beego.AppConfig.String("mysqluser")+
			":"+
			beego.AppConfig.String("mysqlpass")+
			"@tcp(115.159.196.171:3306)/db_name?charset=utf8")
	// register model
	orm.RegisterModel(new(Users))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := models.User{
		User_name:"loongshawn",
		Age:29,
		Email:"loongshawn@jjj.com",
		Staff_id:"123456",
		Position:"研发工程师",
		Extension_number:"",
		Telephone_number:"15642314234",
		Office_location:"商会产业园",
	}

	// insert
	o.Begin()
	id, err := models.UserAdd(o,user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err != nil {
		o.Rollback()
		fmt.Println("插入user表出错,事务回滚")
	} else {
		o.Commit()
		fmt.Println("插入user表成功,事务提交")
	}
}


