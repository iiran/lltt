package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iiran/lltt/pkg/logger"
	"github.com/iiran/lltt/pkg/service/simple_store"
)

import "github.com/iiran/lltt/pkg/setting"

import "github.com/iiran/lltt/pkg/db"

import "github.com/iiran/lltt/pkg/router"

func main() {
	var (
		err        error
		configPath = "config.json"
	)
	setting.Setup(configPath)
	db.Setup(setting.GetDBConfig())
	logger.Setup(setting.GetLogConfig())
	simple_store.Setup(setting.GetSessionConfig())
	gin.SetMode(setting.GetMode())
	r := router.Init()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	logger.Info("server started...")
	if err = r.Run(fmt.Sprintf(":%d", setting.Cfg.Port)); err != nil {
		logger.Error("server stopped... %s", err.Error())
	}
}
