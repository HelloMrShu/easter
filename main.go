package main

import (
	. "code.sohuno.com/sky/robin/global"
	"code.sohuno.com/sky/robin/initialize"
	"code.sohuno.com/sky/robin/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	initialize.Init()
	engine := gin.New()
	router.InitRouter(engine)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Error("服务启动失败", zap.String("error ", err.Error()))
		}
	}()

	// 关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("开始关闭服务器 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("关闭服务器异常:", err)
	}

	fmt.Println("服务器已关闭")
}
