package controller

import (
	"app.com/v1/goadmin/service"
	_ "app.com/v1/goadmin/service"
	"app.com/v1/goadmin/util"
	"app.com/v1/goadmin/validates"
	_ "app.com/v1/goadmin/validates"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris"
)


//登陸接口

func UserLogin(ctx iris.Context) {
	aul := new(validates.LoginRequest)
	fmt.Println(aul.Username)

	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}

	err := validates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(ApiResource(false, nil, e))
				return
			}
		}
	}

	//fmt.Println(aul.Username)

	ctx.Application().Logger().Infof("%s 登录系统", aul.Username)
	ctx.StatusCode(iris.StatusOK)

	admin, exist := service.GetByAdminNameAndPassword(aul.Username, aul.Password)
	if exist {


		session := Sess.Start(ctx)

		//session.Set("adminId", 1)
		//session.Set("isLogin",true)
		//管理员存在 设置session
		//userByte, _ := json.Marshal(admin)

		session.Set("admin", admin.AdminId)
		//redis := datasource.NewRedis()
		//设置session的同步位置为redis

		//session.Set(ISLOGIN, true)

		//fmt.Println(admin)



		_, _ = ctx.JSON(Api(1, "登录成功", "管理员登录成功"))
		return
	}else {
		_, _ = ctx.JSON(Api(0, "登陸失敗", "用戶名或密碼錯誤"))
		return
	}
}



func GetInfo(ctx iris.Context) {
	session := Sess.Start(ctx)
	userByte := session.Get("admin")
	//fmt.Println(userByte)
	if userByte==nil {
		_, _ = ctx.JSON(Api(utils.RECODE_UNLOGIN, utils.EEROR_UNLOGIN, utils.Recode2Text(utils.EEROR_UNLOGIN)))
		return
	}
	adminId, err := session.GetInt64("admin")
	if err!=nil {
		_, _ = ctx.JSON(Api(utils.RECODE_UNLOGIN, utils.EEROR_UNLOGIN, utils.Recode2Text(utils.EEROR_UNLOGIN)))
		return
	}

	id, b := service.GetByAdminId(adminId)
	if !b {
		_, _ = ctx.JSON(Api(0, "登录失败", "用户名或者密码错误,请重新登录"))
		return
	}
	_, _ = ctx.JSON(Api(1, id.AdminToRespDesc(),""))

	return

}

//退出

func GetSingout(ctx iris.Context) {

	//删除session，下次需要从新登录
	session := Sess.Start(ctx)
	session.Delete("admin")

	_, _ = ctx.JSON(Api(utils.RECODE_OK, nil,utils.Recode2Text(utils.RESPMSG_SIGNOUT)))

	return
}