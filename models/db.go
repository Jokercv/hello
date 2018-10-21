package models

import ("time"
		"github.com/astaxie/beego"
		"github.com/astaxie/beego/orm"
		_"github.com/go-sql-driver/mysql"
)

func init(){
	var dbhost string
	var dbport string
	var dbuser string
	var dbpassword string
	var db string
	//获取配置文件中对应的配置信息
	dbhost = beego.AppConfig.String("dbhost")
	dbport = beego.AppConfig.String("dbport")
	dbuser = beego.AppConfig.String("dbuser")
	dbpassword = beego.AppConfig.String("dbpassword")
	db = beego.AppConfig.String("db")
	orm.RegisterDriver( "mysql", orm.DRMySQL) //注册mysql Driver
	//构造conn连接
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8&loc=Local"
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)

	orm.RegisterModel(new(UserInfo),new(RoleInfo))//注册模型
	orm.RunSyncdb("default", false, true)

}
type UserInfo struct {
	Id int
	UserName string
	UserPwd string
	Remark string
	AddDate time.Time
	ModifDate time.Time
	DelFlag int//删除标记
	Roles []*RoleInfo`orm:"rel(m2m)"`

}
type RoleInfo struct {
	Id int
	RoleName string`orm:"size(32)"`
	Remark string
	DelFlag int
	AddData time.Time
	ModifData time.Time
	Users []*UserInfo`orm:"reverse(many)"`
}