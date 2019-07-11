package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer


type User struct {
	Id 			int64		`orm:"auto;PK"`
	Email 		string
	Uid 		string
	Pword		string
	Active		int			`orm:"default(0)"`	  //0:正常，1:禁用
	Profile		*Profile 	`orm:"rel(one)"`
	Role		*Role	 	`orm:"rel(one)"`
	UsersLog	[]*UsersLog `orm:"reverse(many)"`
	CreateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}

type Role struct {
	Id 			int64		`orm:"auto;PK"`
	Name		string
	Status		int			`orm:"default(0)"` 	 // 0:普通用户,1:管理员,2:超级管理员
	Description string		`orm:"default(null)"`
}

type Profile struct {
	Id 			int64		`orm:"auto;PK"`
	User 		*User		`orm:"reverse(one)"`
	Nickname	string		`orm:"default(null)"`
	Age			int			`orm:"default(0)"`   // 0 is nil
	Sex 		int 		`orm:"default(0)"`   //male:1, female:2
	Phone 		string		`orm:"default(null)"`
	WeixinId	string		`orm:"default(null)"`
	Beans		int			`orm:"default(0)"`
	Icon		string		`orm:"default(null)"`
	Photo		string		`orm:"default(null)"`
	LoginCount	int			`orm:"default(0)"`
	UpdateTime	time.Time	`orm:"auto_now;type(datetime)`
}

type UsersLog struct {
	Id 			int64		`orm:"auto;PK"`
	User 		*User		`orm:"rel(fk)"`
	LoginIp		string
	LoginTime	time.Time	`orm:"auto_now_add;type(datetime)`
	Device		string		`orm:"default(null)"`
}

type RushOrder struct {
	Id 			int64		`orm:"auto;PK"`
	RushId		string
	UserId		int64
	UserUid		string
	RushNum		int
	CreateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}

type ConstValue struct {
	Id 			int64		`orm:"auto;PK"`
	ConstType	int
	ValueStr	string
	UseType		int			`orm:"default(1)"`
	UpdateTime	time.Time	`orm:"auto_now;type(datetime)`
	CreateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}


func init() {
	orm.RegisterModel(new(User), new(Profile),
		new(Role), new(ConstValue), new(RushOrder),
		new(UsersLog))
}


func Connect() {
	db_host := beego.AppConfig.String("database::db_host")
	db_type := beego.AppConfig.String("database::db_type")
	db_port := beego.AppConfig.String("database::db_port")
	db_user := beego.AppConfig.String("database::db_user")
	db_pass := beego.AppConfig.String("database::db_pass")
	db_name := beego.AppConfig.String("database::db_name")

	//db_sslmode := beego.AppConfig.String("db_sslmode")
	fmt.Println("data_info:", db_type, db_name, db_host, db_port, db_user, db_pass)
	var dns string

	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", db_user, db_pass, db_host, db_port, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed.", db_type)
		break
	}

	orm.RegisterDataBase("default", db_type, dns)
}

func InsertSuper() {
	r := Role{Status: 2}
	o.Read(&r, "Status")
	//var role
	//qs := o.QueryTable("role")
	//qs.Filter("role__status", 2).Filter("role__name", "SuperAdmin").Values(&role)
	fmt.Println("roles:", r.Id, r.Status, r.Name, r.Description)
	user := new(User)
	user.Active = 0
	user.Email = "123123@admin.com"
	user.Pword = "123123"
	user.Uid = "100011"
	user.Role = &r

	profile := new(Profile)
	profile.UpdateTime = time.Now()
	_, err := o.Insert(profile)
	if err != nil {
		fmt.Println(err)
	}
	user.Profile = profile
	user.CreateTime = time.Now()
	_, err = o.Insert(user)
	if err != nil {
		fmt.Println(err)
	}

	//m2m := o.QueryM2M(&user, "Role")
	//n, err := m2m.Add(&r)
	//if err != nil {
	//	fmt.Println("failed..")
	//}
	//return n
}

func InsertRoles() {
	fmt.Println("insert roles...")
	role := [3]Role{
		{Name: "User", Status: 0, Description: "normal user"},
		{Name: "Admin", Status: 1, Description: "admin"},
		{Name: "SuperAdmin", Status: 2, Description: "super user"},
	}
	for _, v := range role {
		r := new(Role)
		r.Name = v.Name
		r.Status = v.Status
		r.Description = v.Description
		o.Insert(r)
	}
}

func InserConst() {
	fmt.Println("insert const value...")
	cv := [2]ConstValue{
		{ConstType: 1, ValueStr: "100067", UpdateTime: time.Now(), CreateTime: time.Now()},
		{ConstType: 2, ValueStr: "100067", UpdateTime: time.Now(), CreateTime: time.Now()},
	}
	for _, v := range cv {
		c := new(ConstValue)
		c.CreateTime = v.CreateTime
		c.UpdateTime = v.UpdateTime
		c.ConstType = v.ConstType
		c.ValueStr = v.ValueStr

		o.Insert(c)
	}
}

//func Createdb() {
//	db_host := beego.AppConfig.String("db_host")
//	db_type := beego.AppConfig.String("db_type")
//	db_port := beego.AppConfig.String("db_port")
//	db_user := beego.AppConfig.String("db_user")
//	db_pass := beego.AppConfig.String("db_pass")
//	db_name := beego.AppConfig.String("db_name")
//
//}

func Syncdb() {
	Connect()
	o = orm.NewOrm()
	name := "default"
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	InsertRoles()
	InsertSuper()
	InserConst()

	fmt.Println("Database init is complete.")
}