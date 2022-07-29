package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mohwa/ci-cd-github-action/internal/db"
	"github.com/mohwa/ci-cd-github-action/internal/protocol"
)

var (
	debugMode = true
)

func main() {
	defer close()

	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	corsConf := cors.DefaultConfig()
	// Todo: debug 모드 아닐 시에는 tower client만 allow origins에 추가
	if debugMode {
		corsConf.AllowAllOrigins = true
	}
	// middleware 사용 예제
	// https://github.com/gin-gonic/gin#using-middleware
	router := gin.Default()

	protocol.InitRouterGroupRest(router.Group("/api/rest"))
	protocol.InitRouterGroupGraphQL(router.Group("/api/graphql"))

	server := &http.Server{
		Addr:           "127.0.0.1:3001",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.SetKeepAlivesEnabled(false)

	if err := packageInit(); err != nil {
		return
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// signal channel 을 생성한다.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func close() {
	db.Close()
}

func packageInit() error {
	// https://gorm.io/docs/connecting_to_the_database.html

	if err := db.Init(fmt.Sprintf("%s?%s", "yanione:password@tcp(127.0.0.1:3310)/tower", "charset=utf8mb4&parseTime=True&loc=Local")); err != nil {
		return err
	}

	return nil
}
