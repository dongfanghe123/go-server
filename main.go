//package main
//
//import (
//	"go-server/dao"
//	"go-server/handler"
//	"go-server/logger"
//	"go-server/manager"
//	"go-server/router"
//	"go-server/service"
//	"go-server/util"
//	"log"
//	"net/http"
//	_ "net/http/pprof"
//
//	_ "github.com/go-sql-driver/mysql" // 必须引入数据库驱动
//)
//
//func main() {
//
//	logger.InitLogger()
//
//	if err := dao.Init("resources/juice.xml"); err != nil {
//		log.Fatal(err)
//	}
//	defer dao.Close()
//
//	dao.InitDao()
//	manager.Init()
//	util.Init(1)
//
//	voucherDao := dao.NewVoucherDao()
//	voucherService := service.NewVoucherService(voucherDao)
//	voucherHandler := handler.NewVoucherHandler(voucherService)
//
//	myHandler := &handler.Handler{
//		VoucherHandler: voucherHandler,
//	}
//
//	r := router.NewRouter(myHandler)
//
//	go func() {
//		http.ListenAndServe(":8848", nil)
//	}()
//
//	logger.Log.Info().Msg("app start success")
//
//	r.Run(":8081")
//
//}

package main

import (
	"fmt"
	"time"
)

func foo() (result int) {
	defer func() {
		result++
	}()

	return 10
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		fmt.Println("send done")
	}()

	fmt.Println(<-ch)
	time.Sleep(time.Second)
}
