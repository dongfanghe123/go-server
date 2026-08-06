package main

import (
	"go-server/dao"
	"go-server/manager"
	"go-server/router"
	"go-server/util"
	"log"

	_ "github.com/go-sql-driver/mysql" // 必须引入数据库驱动
)

func main() {

	if err := dao.Init("resources/juice.xml"); err != nil {
		log.Fatal(err)
	}
	defer dao.Close()

	dao.InitDao()
	manager.Init()
	util.Init(1)

	r := router.NewRouter()
	r.Run(":8081")

}
