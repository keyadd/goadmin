package main

import (
	"app.com/v1/goadmin/config"
	"app.com/v1/goadmin/controller"
	"app.com/v1/goadmin/datasource"
	_ "app.com/v1/goadmin/datasource"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

func main()  {

	app :=  newApp()

	//app設置
	configation(app)

	//路由設置
	mavcHandle(app)
	datasource.NewMysqlEngine()

	//启用session




	config := config.InitConfig()
	addr := ":"+ config.Port
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}



var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
func mavcHandle(app *iris.Application)  {
	api :=app.Party("/").AllowMethods(iris.MethodOptions)
	{
		v1 := api.Party("admin/")
		{
			v1.Post("login",controller.UserLogin)
			v1.Get("info",controller.GetInfo)
			v1.Get("singout",controller.GetSingout)
		}

	}


}

func Session(ctx iris.Context) {
	//除了登录接口以外,其他接口都需要进行session验证
	if ctx.Path() != "/admin/login" {
		// 检查用户是否已通过身份验证
		s :=  controller.Sess.Start(ctx).Get("zzy")
		if s == nil {
			ctx.StatusCode(512)
			return
		}
		//获取请求头里的session,如果与内置的session一致则通过校验
		e := ctx.Request().Header.Get("Cookie")
		if s == e == false {
			ctx.StatusCode(512)
			return
		}
	}
	ctx.Next()
}




func basicMVC(app *mvc.Application) {
	//当然，您可以在MVC应用程序中使用普通的中间件.
	app.Router.Use(func(ctx iris.Context) {
		//sessManager := sessions.New(sessions.Config{
		//	Cookie:  "sessioncookie",
		//	Expires: 24 * time.Hour,
		//})
		////start := sessManager.Start​,
		////isLogin, err := sessManager.GetBoolean(ISLOGIN)
		//if start != nil {
		//	return
		//}
		ctx.Next()
	})
}


func newApp() *iris.Application {
	app := iris.New()
	//設置日誌級別
	app.Logger().SetLevel("debug")
	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manage/static","./static")

	//z註冊視圖模板文件
	app.RegisterView(iris.HTML("./static",".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})
	return app
}




func configation(app *iris.Application)  {
	//配置字符編碼
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset:"UTF-8",
	}))


	//配置錯誤提示
	app.OnErrorCode(iris.StatusNotFound, func(c iris.Context) {
		c.JSON(iris.Map{
			"errmsg":iris.StatusNotFound,
			"msg":"interal error",
			"data": iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(c iris.Context) {
		c.JSON(iris.Map{
			"errmsg":iris.StatusInternalServerError,
			"msg":"interal error",
			"data": iris.Map{},
		})
	})
}