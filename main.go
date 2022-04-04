package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"klsb/common"
	"klsb/router"
)

func main() {
	r := gin.Default()
	common.InitConfig("config/config.yaml")
	common.InitDB()

	//创建DB实例进行数据库操作
	db := common.GetDB()
	//延迟关闭数据库
	defer db.Close()

	//r.Use(middleware.Cors())
	r = router.InitRouter(r)

	addr := common.Conf.APP.Addr
	port := common.Conf.APP.Port

	r.Run(fmt.Sprintf("%s:%s", addr, port))
}
