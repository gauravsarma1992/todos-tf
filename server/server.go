package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	ApiServer struct {
		engine *gin.Engine
		config Config
		db     *gorm.DB
	}

	Config struct {
		ServerHost string `json:"server_host"`
		ServerPort string `json:"server_port"`
		DBHost     string `json:"db_host"`
		DBPort     string `json:"db_port"`
		DBUsername string `json:"db_username"`
		DBPassword string `json:"db_password"`
		DBName     string `json:"db_name"`
	}
)

func New(config Config) (apiServer *ApiServer, err error) {
	apiServer = &ApiServer{
		config: config,
		engine: gin.Default(),
	}
	if err = apiServer.setupRoutes(); err != nil {
		return
	}
	if err = apiServer.dbConnect(); err != nil {
		return
	}
	if err = apiServer.dbMigrate(); err != nil {
		return
	}
	return
}

func getDefaultConfig() (config Config) {
	config = Config{
		ServerHost: "127.0.0.1",
		ServerPort: "8090",
		DBHost:     "127.0.0.1",
		DBPort:     "3306",
		DBUsername: "root",
		DBPassword: "",
		DBName:     "todo_tf",
	}
	return
}

func Default() (apiServer *ApiServer, err error) {
	if apiServer, err = New(getDefaultConfig()); err != nil {
		return
	}
	return
}

func (apiServer *ApiServer) setupRoutes() (err error) {
	apiServer.engine.GET("/todos/:id", apiServer.TodoShow)
	apiServer.engine.POST("/todos", apiServer.TodoCreate)
	return
}
func (apiServer *ApiServer) dbMigrate() (err error) {
	apiServer.db.AutoMigrate(&TodoItem{})
	return
}

func (apiServer *ApiServer) dbConnect() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		apiServer.config.DBUsername,
		apiServer.config.DBPassword,
		apiServer.config.DBHost,
		apiServer.config.DBPort,
		apiServer.config.DBName,
	)

	if apiServer.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return
	}
	return
}

func (apiServer *ApiServer) Run() (err error) {
	err = apiServer.engine.Run(fmt.Sprintf("%s:%s", apiServer.config.ServerHost, apiServer.config.ServerPort))
	return
}
