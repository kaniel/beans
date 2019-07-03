package models

import (
	"time"
)

type User struct {
	Id 			int			`orm:"auto;PK"`
	email 		string
	Uid 		string
	Pword		string
	Active		int			`orm:"default:0"`  //0:正常，1:禁用
	UserInfo	*UserInfo 	`orm:"rel(one)"`
	UserLog		*UserLog 	`orm:"reverse(many)"`
	CreateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}

type Role struct {
	Id 			int			`orm:"auto;PK"`
	Name		string
	Mark		int			`orm:"default:0"`  // 0:普通用户,1:管理员,2:超级管理员
	Description string
}

type UserInfo struct {
	Id 			int			`orm:"auto;PK"`
	UserId 		int
	Nickname	string
	Age			int			`orm:"default:0"`  // 0 is nil
	sex 		int 		`orm:"default:0"`  //male:1, female:2
	Phone 		string
	WeixinId	string
	Beans		int			`orm:"default:0"`
	Icon		string
	Photo		string
	LoginCount	int			`orm:"default:0"`
	UpdateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}

type UserLog struct {
	Id 			int			`orm:"auto;PK"`
	loginIp		string
	loginTime	time.Time	`orm:"auto_now_add;type(datetime)`
	device		string
}

type RushOrder struct {
	Id 			int			`orm:"auto;PK"`
	RushId		string
	UserId		int
	UserUid		string
	RushNum		int			`orm:"default:0"`
	CreateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}

type ConstValue struct {
	Id 			int			`orm:"auto;PK"`
	ConstType	int
	ValueStr	string
	UseType		int			`orm:"default:0"`
	UpdateTime	time.Time
	CreateTime	time.Time	`orm:"auto_now_add;type(datetime)`
}