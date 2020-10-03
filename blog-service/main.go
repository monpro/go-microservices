package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/pkg/settting"
	"github.com/go-microservices/blog-service/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	server := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	// log the settings
	//fmt.Printf("%+v\n", global.ServerSetting)
	//fmt.Printf("%+v\n", global.AppSetting)
	//fmt.Printf("%+v\n", global.DatabaseSetting)
	server.ListenAndServe()
}

func setupSetting() error {
	newSetting, err := settting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil

}
