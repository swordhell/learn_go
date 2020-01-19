package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 必须引入，如若不引入，则需要自己定义
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // 必须引入，不然报错
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	eng := engine.Default()
	cfg := config.Config{}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// 从.so文件中加载插件
	examplePlugin := plugins.LoadFromPlugin("../datamodel/example.so")

	eng.AddConfig(cfg).
		AddPlugins(adminPlugin, examplePlugin). // 加载插件
		Use(r)

	r.Run(":9033")
}
