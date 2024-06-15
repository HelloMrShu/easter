package main

import (
	"context"
	"errors"
	"fmt"
	. "github.com/HelloMrShu/easter/global"
	"github.com/HelloMrShu/easter/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// init
	Initialize()

	// load router
	engine := gin.New()
	engine.Use(LoggerMiddleware(Logger), RecoveryMiddleware(Logger, false))
	router.InitRouter(engine)

	// start
	srv := &http.Server{
		Addr:    ":8080",
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
		log.Fatal("shutdown failed:", err)
	}

	fmt.Println("shutdown success")
}
