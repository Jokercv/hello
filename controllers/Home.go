package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"github.com/astaxie/beego/orm"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Index() {
	this.TplName="Home/Index.html"
}
func (this *HomeController) ShowIndex() {
	this.TplName = "Home/ShowIndex.html"
}
func (this *HomeController)GetMenus()  {
	userId:=this.GetSession("userId")
	var userInfo models.UserInfo
	o:=orm.NewOrm()
	o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
	var roles []*models.RoleInfo
	o.LoadRelated(&userInfo,"Roles")
	for _,role:=range userInfo.Roles{
		roles=append(roles,role)

	}
	//var actions []models.ActionInfo
	//for i:=0	;i<len(roles) ;i++  {
	//	o.LoadRelated(roles[i],"Actions")
	//
	//	for _,action:=range roles[i].Actions{
	//		actions=append(actions,action)
	//	}
	//}
	//var menuActions []*models.ActionInfo
	//for i:=0;i<len(actions);i++{
	//	if actions[i].ActionTypeEnum == 1 {
	//		menuActions=append(menuActions,actions[i])
	//	}
	//}
}

