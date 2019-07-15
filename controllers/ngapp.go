package controllers


func (ng *NgAppController) Get() {
	ng.Data["Website"] = "beego.me"
	ng.Data["Email"] = "astaxie@gmail.com"
	ng.TplName = "ng.tpl"
}

func (ng *NgAppController) Index() {
	ng.Data["Website"] = "beego.me"
	ng.Data["Email"] = "astaxie@gmail.com"
	ng.TplName = "index.tpl"
}
