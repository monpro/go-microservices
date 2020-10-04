package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-microservices/blog-service/global"
	"github.com/go-microservices/blog-service/internal/model"
	"github.com/go-microservices/blog-service/internal/routers"
	"github.com/go-microservices/blog-service/pkg/logger"
	"github.com/go-microservices/blog-service/pkg/settting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("inint setupLogger: %v", err)
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

	global.Logger.InfoFormat("%s: go/%s","test", "blog-service")
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

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDbEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +
			global.AppSetting.LogFileName +
			global.AppSetting.LogFileExt,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}