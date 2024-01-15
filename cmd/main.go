package main

import (
	"fmt"

	"github.com/Asliddin3/tz/config"
	"github.com/Asliddin3/tz/controller"
	"github.com/Asliddin3/tz/docs"
	"github.com/Asliddin3/tz/migrate"
	"github.com/Asliddin3/tz/pkg/logger"
	"github.com/Asliddin3/tz/pkg/middleware"
	postgresdb "github.com/Asliddin3/tz/pkg/postgres"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go.uber.org/automaxprocs"
)

// @title          Golan tz documentation
// @version         1.0
// @description     This is a sample server CRM server.

// @securityDefinitions.apikey  BearerAuth
// @host localhost:8000
// @in header
// @name Authorization
// @Description									AUTH.
func main() {
	cfg := config.Load()
	log := logger.NewLogger()
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.json", // "./openapi.yaml"
		SpecPath:    "/docs/swagger.json",  // "
		DocsPath:    "/redoc",
	}

	server := gin.Default()
	server.Use(
		ginredoc.New(doc),
		gin.Recovery(),
		gin.Logger(),
		middleware.New(middleware.GinCorsMiddleware()),
		gzip.Gzip(gzip.DefaultCompression),
	)

	db, err := postgresdb.NewClient(cfg)
	if err != nil {
		log.Error("postgresdb connection error", logger.Error(err))
		return
	}
	log.Debug("postgresql connection established")

	err = migrate.Migrate(db)
	if err != nil {
		log.Error("failed to migrate db", logger.Error(err))
		return
	}
	log.Debug("migrations complete")

	handler := controller.NewHandler(db, log, cfg)

	handler.Init(server)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = server.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Error("router running error", logger.Error(err))
		return
	}
}
