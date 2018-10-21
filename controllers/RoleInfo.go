package controllers

import (
	"github.com/astaxie/beego"
	"hello/models"
	"time"
	"github.com/astaxie/beego/orm"
)

type RoleInfoController struct {
	beego.Controller
}

func (this *RoleInfoController) Index() {
	this.TplName="RoleInfo/Index.html"
}
func (this *RoleInfoController) ShowAddRole() {
	this.TplName="RoleInfo/ShowAddRole.html"
}
func (this *RoleInfoController) AddRole() {
	var roleInfo  =models.RoleInfo{}
	roleInfo.RoleName=this.GetString("roleName")
	roleInfo.Remark=this.GetString("roleRemark")
	roleInfo.DelFlag=0
	roleInfo.AddData=time.Now()
	roleInfo.ModifData=time.Now()
	//beego.Info(roleInfo,"----------------")
	o:=orm.NewOrm()
	_,err:=o.Insert(&roleInfo)
	if err != nil {
		this.Data["json"]=map[string]interface{}{"flag":"no"}
		beego.Info("errrrrrrrrr",err)
	}else {this.Data["json"]=map[string]interface{}{"flag":"ok"}}
	this.ServeJSON(
	)
}
//获取角色信息：
func (this *RoleInfoController)GetRoleInfo(	)  {
	pageIndex,_:=this.GetInt("page")
	pageSize,_:=this.GetInt("rows")
	start:=(pageIndex-1)*pageSize
	o:=orm.NewOrm()
	var roles []models.RoleInfo
	o.QueryTable("role_info").Filter("del_flag",0).OrderBy("Id").Limit(pageSize,start).All(&roles)
	count,err:= o.QueryTable("role_info").Filter("del_flag",0).Count()
	beego.Info("countttttttt",count)
	if err != nil {
		beego.Info("errxxxxx",err)
		return
	}
	this.Data["json"]=map[string]interface{}{"rows":roles,"total":count}
	this.ServeJSON()
}