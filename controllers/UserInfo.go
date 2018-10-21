package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"hello/models"
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
)

type UserInfoController struct {
	beego.Controller
}
type UserData struct {
	PageIndex int
	PageSize int
	Name string
	Remark string
	Totalcount int64
}
func (this *UserInfoController)Index()  {
this.TplName="UserInfo/Index.html"
}
func (this *UserInfoController) GetUserInfo() {
	//1:接收传递过来的当前页码
	pageIndex,_:=strconv.Atoi(this.GetString("page"))
	//2；接收传递过来的每页记录数
	pageSize,_:=strconv.Atoi(this.GetString("rows"))
	name:=this.GetString("name")
	remark:=this.GetString("remark")
	var userSearchData  =UserData{}
	userSearchData.Remark=remark
	userSearchData.Name=name
	userSearchData.PageSize=pageSize
	userSearchData.PageIndex=pageIndex
	serverData:=userSearchData.SearchUserData()
	this.Data["json"]=map[string]interface{}{"rows":serverData,"total":userSearchData.Totalcount}

	//start:=(pageIndex-1)*pageSize
	//var users []models.UserInfo
	//o:=orm.NewOrm()
	//o.QueryTable("user_info").Filter("del_flag",0).OrderBy("id").Limit(pageSize,start).All(&users)
	//count,_:=o.QueryTable("user_info").Filter("del_flag",0).Count()
	//this.Data["json"]=map[string]interface{}{"rows":users,"total":count}
	this.ServeJSON()
}
func (this *UserData) SearchUserData() []models.UserInfo {
	o:=orm.NewOrm()
	temp:=o.QueryTable("user_info")
	if this.Name!="" {
		//__icontains不区分大小写筛选
		temp=temp.Filter("user_name__icontains",this.Name)
	}
	if this.Remark!="" {
		temp=temp.Filter("remark__icontains",this.Remark)
	}
	temp=temp.Filter("del_flag",0)
	var err error
	this.Totalcount,err=temp.Count()
	if err != nil {
		beego.Info("err=", err)
		return nil
	}
	start:=(this.PageIndex-1)*this.PageSize
	var users []models.UserInfo
	temp.OrderBy("Id").Limit(this.PageSize,start ).All(&users)

	return users
	//temp=temp.Filter("del_flag",0)
	//this.Totalcount,_=temp.Count()
	//start:=(this.PageSize-1)*this.PageSize
	//var users[]models.UserInfo
	//temp.OrderBy("Id").Limit(this.PageSize,start).All(&users)
	//return users
}
func (this *UserInfoController) AddUser() {
	var userInfo  =models.UserInfo{}
	userInfo.UserName=this.GetString("userName")
	userInfo.UserPwd=this.GetString("userPwd")
	userInfo.Remark=this.GetString("userRemark")

	userInfo.ModifDate=time.Now()
	userInfo.AddDate=time.Now()
	//beego.Info(userInfo)
	userInfo.DelFlag=0//表示正常，1表示表示软删除。
	o:=orm.NewOrm()
	_,err:=o.Insert(&userInfo)
	if err == nil {
		this.Data["json"]=map[string]interface{}{"flag":"ok"}

	}else {
		this.Data["json"]=map[string]interface{}{"flag":"no"}
		//beego.Info(err)
	}
	this.ServeJSON()
}
func (this *UserInfoController)DeleteUser()  {
	ids:=this.GetString("strId")
	strIds:=strings.Split(ids,"-")
	o:=orm.NewOrm()
	var userInfo=models.UserInfo{}
	for i:=0;i<len(strIds);i++ {
		id,_:=strconv.Atoi(strIds[i])
		userInfo.Id=id
		o.Delete(&userInfo)
	}
	this.Data["json"]=map[string]interface{}{"flag":"ok"}
	this.ServeJSON()

}
func (this *UserInfoController) ShowEditUser() {
	userId,_:=this.GetInt("userId")
	var userInfo models.UserInfo
	o:=orm.NewOrm()
	o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
	this.Data["json"]=map[string]interface{}{"userInfo":userInfo}
	this.ServeJSON()

}
func (this *UserInfoController) EditUser() {
	var userInfo  =models.UserInfo{}
	userInfo.UserName=this.GetString("userEditName")
	userInfo.UserPwd=this.GetString("userEditPwd")
	userInfo.Remark=this.GetString("userEditRemark")
	userInfo.Id,_=this.GetInt("userEditId")
	//this.Data["json"]=map[string]interface{}{}
	o:=orm.NewOrm()
	_,err:=o.Update(&userInfo)
	if err != nil {
		this.Data["json"]=map[string]interface{}{"flag":"ok"}
		beego.Info("errrrr",err)
	}else {
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}

	this.ServeJSON()

}
func (this *UserInfoController)ShowSetUserRole()  {
	userId,_:=this.GetInt("userId")
	var userInfo models.UserInfo
	o:=orm.NewOrm()
	o.QueryTable("user_info").Filter("id",userId).One(&userInfo)
	var userExtRoles []*models.RoleInfo//用户已有角色
	o.LoadRelated(&userInfo,"Roles")
	for _,role:=range userInfo.Roles{
		userExtRoles=append(userExtRoles,role)
	}
	//查询所有角色
	var allRoles[]models.RoleInfo
	o.QueryTable("role_info").Filter("del_flag",0).All(&allRoles)
	this.Data["allRoles"]=allRoles
	this.Data["userExtRoles"]=userExtRoles
	this.Data["userInfo"]=userInfo
	this.TplName="UserInfo/ShowSetUserRole.html"
}
