package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	. "github.com/HelloMrShu/easter/global"
	"github.com/HelloMrShu/easter/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	env  string
	port int
)

func main() {
	flag.StringVar(&env, "e", "dev", "env")
	flag.IntVar(&port, "p", 8080, "port")

	// initialize
	Initialize(env)

	// load router
	engine := gin.New()
	engine.Use(LoggerMiddleware(Logger), RecoveryMiddleware(Logger, false))
	router.InitRouter(engine)

	// start
	addr := fmt.Sprintf("%s:%v", "", port)
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			Logger.Error("server start failed", zap.String("error ", err.Error()))
		}
	}()

	// shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		Logger.Fatal("shutdown failed:" + err.Error())
	}
}
