package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/Admin/UserInfo/Index",&controllers.UserInfoController{},"get:Index")
    beego.Router("/Admin/UserInfo/AddUser",&controllers.UserInfoController{},"post:AddUser")
    beego.Router("/Admin/UserInfo/GetUserInfo",&controllers.UserInfoController{},"post:GetUserInfo")
    beego.Router("/Admin/UserInfo/DeleteUser",&controllers.UserInfoController{},"post:DeleteUser")
    beego.Router("/Admin/UserInfo/ShowEditUser",&controllers.UserInfoController{},"post:ShowEditUser")
    beego.Router("/Admin/UserInfo/EditUser",&controllers.UserInfoController{},"post:EditUser")
    beego.Router("/Admin/UserInfo/ShowSetUserRole",&controllers.UserInfoController{},"get:ShowSetUserRole")
    //=--------------------------------RoleInfo----------------------
    beego.Router("/Admin/RoleInfo/Index",&controllers.RoleInfoController{},"get:Index")
    beego.Router("/Admin/RoleInfo/ShowAddRole",&controllers.RoleInfoController{},"get:ShowAddRole")
    beego.Router("/Admin/RoleInfo/AddRole",&controllers.RoleInfoController{},"post:AddRole")
    beego.Router("/Admin/RoleInfo/GetRoleInfo",&controllers.RoleInfoController{},"post:GetRoleInfo")
    //beego.Router("/Admin/UserInfo/ShowSetUserAction",&controllers.UserInfoController{},"get:ShowSetUserAction")
    //----------------------------------Home-----------------
    beego.Router("/Admin/Home/ShowIndex",&controllers.HomeController{},"get:ShowIndex")
    beego.Router("/Admin/Home/Index",&controllers.HomeController{},"get:Index")
}
