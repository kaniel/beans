package controllers

import (
	"beans/commons"
	"beans/models"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
)

var o orm.Ormer

func (u *UserController) Get() {

}

func (u *UserController) Register() {


	email := u.GetString("email")
	password := u.GetString("password")

	if len(email) == 0 || commons.VerificationEmail(email) == false {
		u.Data["111"] = &map[string] interface{} {"code":1}
		u.ServeJSON()
	}

	if len(password) < 8 || len(password) > 16 || commons.VerificationPassword(password) == false {
		u.Data["111"] = &map[string] interface{} {"code":2}
		u.ServeJSON()
	}

	r := models.Role{Status:0}
	o.Read(&r, "Status")

	cvalue := models.ConstValue{ConstType:1}
	o.Read(&cvalue, "ConstType")

	user_uid := strconv.Atoi(cvalue.ValueStr)


	user := new(models.User)
	user.Email = email
	user.Password = password
	user.Uid = user_uid
}