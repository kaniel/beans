package controllers

import (
	"github.com/astaxie/beego"
)


type BaseController struct {
	beego.Controller
}

type NgAppController struct {
	BaseController
}

type UserController struct {
	BaseController
}

type AuthController struct {
	BaseController
}