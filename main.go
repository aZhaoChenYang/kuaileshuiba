package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	//r.Use(middleware.Cors())
	r = router.InitRouter(r)

	addr := common.Conf.APP.Addr
	port := common.Conf.APP.Port

	err := r.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		return
	}
}
