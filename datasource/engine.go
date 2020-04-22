package datasource

import (
	"app.com/v1/goadmin/model"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"app.com/v1/goadmin/config"
)

//databases engine

func NewMysqlEngine() *xorm.Engine {

	initConfig := config.InitConfig()

	if initConfig == nil {
		return nil
	}

	database := initConfig.DataBase
	engine, err := xorm.NewEngine("mysql", "root:toor@tcp(127.0.0.1:3306)/goadmin?charset=Utf8")
	iris.New().Logger().Info(database)

	err = engine.Sync2(new(
		model.Permission),
		new(model.City),
		new(model.Admin),
		new(model.User),
		new(model.UserOrder),
		new(model.Address),
		new(model.Shop),
		new(model.OrderStatus),
		new(model.FoodCategory),
		new(model.Food))
	if err != nil {
		panic(err.Error())
	}
	//show sql

	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)
	return engine
}