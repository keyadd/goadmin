package service

import (
	"app.com/v1/goadmin/datasource"
	"app.com/v1/goadmin/model"
)

type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
	GetByAdminId(adminId int64) (model.Admin, bool)
	//获取管理员总数
	GetAdminCount() (int64, error)
	SaveAvatarImg(adminId int64, fileName string) bool
	GetAdminList(offset, limit int) []*model.Admin
}

/**
 * 通过用户名和密码查询管理员
 */
func GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin

	datasource.NewMysqlEngine().Where(" admin_name = ? and pwd = ? ", username, password).Get(&admin)

	return admin, admin.AdminId != 0
}


/**
 * 查询管理员信息
 */
func GetByAdminId(adminId int64) (model.Admin, bool) {
	var admin model.Admin

	datasource.NewMysqlEngine().Id(adminId).Get(&admin)

	return admin, admin.AdminId != 0
}



