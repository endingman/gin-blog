package main

import (
	// fmt：实现了类似C语言printf和scanf的格式化I/O。格式化动作（’verb’）源自C语言但更简单
	//net/http：提供了HTTP客户端和服务端的实现
	"context"
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/gredis"
	"gin-blog/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"time"

	//"github.com/fvbock/endless"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
)

func main() {
	// endless用法
	/**
	  endless.DefaultReadTimeOut = setting.ReadTimeout
	  endless.DefaultWriteTimeOut = setting.WriteTimeout
	  endless.DefaultMaxHeaderBytes = 1 << 20
	  endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	  server := endless.NewServer(endPoint, routers.InitRouter())
	  server.BeforeBegin = func(add string) {
	      log.Printf("Actual pid is %d", syscall.Getpid())
	  }

	  err := server.ListenAndServe()
	  if err != nil {
	      log.Printf("Server err: %v", err)
	  }

	*/

	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	
	//http.Server - Shutdown()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

	//router := routers.InitRouter()

	/**
	  Addr：监听的TCP地址，格式为:8000
	  Handler：http句柄，实质为ServeHTTP，用于处理程序响应HTTP请求
	  TLSConfig：安全传输层协议（TLS）的配置
	  ReadTimeout：允许读取的最大时间
	  ReadHeaderTimeout：允许读取请求头的最大时间
	  WriteTimeout：允许写入的最大时间
	  IdleTimeout：等待的最大时间
	  MaxHeaderBytes：请求头的最大字节数
	  ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
	  ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）
	       * @type {[type]}
	*/
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()

}
