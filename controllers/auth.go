package controllers

import "fmt"

func (au *AuthController) Index() {
	au.TplName = "login.tpl"
}

func (au *AuthController) Login(){
	username := au.GetString("username")
	pwd := au.GetString("password")
	fmt.Println(username, pwd)
	//au.Data["json"] = &map[string] interface{}{"total": 1}
	//au.ServeJSON()
	au.TplName = "index.tpl"
}

func (au *AuthController) Logout() {

}

func (au *AuthController) Refresh() {

}