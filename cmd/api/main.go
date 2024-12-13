package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-clean-architecture/docs"
	"go-clean-architecture/internal/domains/user"
	"go-clean-architecture/pkg"
	"gorm.io/gorm"
	"log"

	"github.com/gin-contrib/cors"
)

func InitializeConfig() *pkg.Config {
	pkg.LoadConfig()

	return pkg.GetConfig()
}

func InitializeDB() (db *gorm.DB) {
	db, err := pkg.Postgres()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}

func InitializeServer() *gin.Engine {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()
	customWriter, logger := pkg.CustomLogger()

	server.Use(logger)
	server.Use(gin.LoggerWithWriter(customWriter))

	return server
}

func InitializeUserDomainRouter(db *gorm.DB, routerGroup *gin.RouterGroup) {
	userRepository := user.NewUserRepository(db)
	userUseCase := user.NewUserUseCase(userRepository)
	userHttp := user.NewUserHTTP(userUseCase)

	user.SetupRoutes(routerGroup, userHttp)
}

func InitializeDocsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

//	@title			Go Clean Architecture
//	@version		1.0
//	@description	This is a documentation for a go refreshment projects
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	devin.com
//	@contact.email	devinanugrahp27@gmail.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

//	@host	localhost:8080
// 	@BasePath /api

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	cfg := InitializeConfig()
	server := InitializeServer()
	db := InitializeDB()

	router := server.Group("/api")

	InitializeDocsRouter(router)
	InitializeUserDomainRouter(db, router)

	err := server.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		_ = fmt.Errorf("error starting server: %v", err)
	}
}
