package main

import (
	"fmt"
	"path/filepath"

	"github.com/anyingiit/GoReactResourceManagement/db"
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
	"github.com/anyingiit/GoReactResourceManagement/router"
	"github.com/anyingiit/GoReactResourceManagement/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	globalVars.InitGlobalVars()

	// 获取项目根目录路径, 并设置到globalVars中
	// 打印状态
	fmt.Println("Getting project root path...")
	projectRootPath, err := filepath.Abs(filepath.Dir("./main.go"))
	if err != nil {
		panic(fmt.Errorf("failed to get project root path, %v", err))
	}
	fmt.Println("Project root path is", projectRootPath)

	// 打印状态
	fmt.Println("Reading config file...")
	err = globalVars.ProjectRootPath.Set(projectRootPath)
	if err != nil {
		panic(fmt.Errorf("failed to set project root path, %v", err))
	}

	config, err := utils.ReadConfigFile(filepath.Join(projectRootPath, "config", "config.yml"))
	if err != nil {
		panic(fmt.Errorf("failed to read config file, %v", err))
	}
	err = globalVars.ProjectConfig.Set(config)
	if err != nil {
		panic(fmt.Errorf("failed to set project config, %v", err))
	}

	// 打印状态
	fmt.Println("Initializing database...")
	db, err := db.InitDB(&config.Database)
	if err != nil {
		panic(fmt.Errorf("failed to init db, %v", err))
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 创建gin实例
	// 打印状态
	fmt.Println("Creating gin instance...")
	g := gin.New()

	// 注册中间件
	// 打印状态
	fmt.Println("Registering middleware...")
	g.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())

	//init router
	// 打印状态
	fmt.Println("Initializing router...")
	router.InitRouter(g, db)

	// create server address string, and print it
	serverAddress := fmt.Sprintf("%s:%d", config.Server.LocalIp, config.Server.LocalPort)
	fmt.Println("Server is running on", serverAddress)

	g.Run(serverAddress)
}

//TODO: 首次运行时, 应该执行一些操作
//TODO: globalVars可以继续抽象
