package server

import (
	"fmt"
	"moges/common/config"
	"moges/common/security"
	"moges/domain/model"
	"moges/logger"
	"moges/storage"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Serve loads the application's config file and starts the server
func Serve() {
	// Initialize all
	initConfig()
	storage.InitializeGorm(&config.Config.Database)

	// Auto migrate
	// Note: This just creates column, does not create index, foreign keys, ...
	storage.GormMaster.AutoMigrate(&model.Photo{})

	// Start http server
	serve()

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credential", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// serve creates the service and starts the server
func serve() {
	// Router config
	router := gin.Default()
	router.Use(CORS())
	router.Use(security.AuthMiddleware())
	router.Use(security.PhotoMiddleware())

	v1 := router.Group("")
	{
		RegisterAPI(v1)
	}

	// Gracefully stop application
	gracefulStop()

	// Start server at default port: 8080
	port := viper.Get("server.port")
	address := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	logger.Info(fmt.Sprintf("Start server at port %d", port))
	err := router.Run(address)
	if err != nil {
		panic(err)
	}
}

// gracefulStop create Go routine to listen channel “gracefulStop” for incoming signals.
// The following Go routine will block until it receives signals from OS
func gracefulStop() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logger.Info("Ticker stopped")
		logger.Info("caught sig: %+v", sig)
		logger.Info("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		storage.Close()
		os.Exit(0)
	}()
}

// initConfig ...
func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("/etc/moges")
	viper.AddConfigPath("$HOME/.moges")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(config.Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into config struct, %v", err))
	}
}
