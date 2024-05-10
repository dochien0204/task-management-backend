package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	authHandler "source-base-go/api/handler/auth"
	masterDataHandler "source-base-go/api/handler/master_data"
	projectHandler "source-base-go/api/handler/project"
	roleHandler "source-base-go/api/handler/role"
	taskHandler "source-base-go/api/handler/task"
	userHandler "source-base-go/api/handler/user"
	"source-base-go/api/middleware"
	"source-base-go/config"
	"source-base-go/infrastructure/repository"
	jwtUtil "source-base-go/infrastructure/repository/util"
	masterdata "source-base-go/usecase/master_data"
	"source-base-go/usecase/project"
	"source-base-go/usecase/role"
	"source-base-go/usecase/task"
	"source-base-go/usecase/user"

	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	"golang.org/x/text/language"
)

func init() {
	config.SetConfigFile("config")
	os.Setenv("TZ", "Etc/GMT")
}

func main() {
	envConfig := getConfig()

	//Database Connect
	db, err := repository.ConnectDB(envConfig.Postgres)
	if err != nil {
		log.Println(err)
		return
	}

	//Redis
	redisClient, err := repository.ConnectRedis(envConfig.Redis)
	if err != nil {
		log.Println(err)
		return
	}

	//Start app
	app := gin.New()
	//Cors
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)

	//Verifier
	verifier := jwtUtil.NewJWTVerifier(redisClient)

	//Define Repository
	userRepo := repository.NewUserRepostory(db)
	roleRepo := repository.NewRoleRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	userProjectRoleRepo := repository.NewUserProjectRole(db)
	projectRepo := repository.NewProjectRepository(db)
	statusRepo := repository.NewStatusRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	taskDocumentRepo := repository.NewTaskDocumentRepository(db)
	emailRepo := repository.NewEmailRepository(db)
	cateRepo := repository.NewCategoryRepository(db)
	discussionRepo := repository.NewDiscussionRepository(db)
	activityRepo := repository.NewActivityRepository(db)

	//Define service
	userService := user.NewService(userRepo, roleRepo, userRoleRepo, statusRepo, emailRepo, verifier)
	roleServce := role.NewService(roleRepo)
	projectService := project.NewService(projectRepo, userProjectRoleRepo, roleRepo)
	masterDataService := masterdata.NewService(statusRepo, cateRepo)
	taskService := task.NewService(taskRepo, statusRepo, taskDocumentRepo, discussionRepo, activityRepo, userRepo, projectRepo)

	//gin I18n
	app.Use(ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         "./resource/lang",
			AcceptLanguage:   []language.Tag{language.English, language.Vietnamese},
			DefaultLanguage:  language.English,
			UnmarshalFunc:    json.Unmarshal,
			FormatBundleFile: "json",
		}),
		ginI18n.WithGetLngHandle(func(ctx *gin.Context, defaultLanguage string) string {
			language := ctx.Query("language")
			if language != "" {
				return language
			}

			return defaultLanguage
		}),
	))

	//Transaction middleware
	tx := middleware.NewMiddlewareRepository(db)

	//Handler
	userHandler.MakeHandlers(app, userService, verifier, tx)
	roleHandler.MakeHandlers(app, roleServce, verifier, tx)
	authHandler.MakeHandlers(app, userService, tx)
	projectHandler.MakeHandlers(app, projectService, verifier, tx)
	masterDataHandler.MakeHandlers(app, masterDataService, verifier)
	taskHandler.MakeHandlers(app, taskService, verifier, tx)

	//Swagger
	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(fmt.Sprintf("%s%s%v", envConfig.Host, ":", envConfig.Port))

}

func getConfig() config.EnvConfig {
	return config.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Postgres: config.PostgresConfig{
			Timeout:  config.GetInt("database.postgres.timeout"),
			DBName:   config.GetString("database.postgres.dbname"),
			Username: config.GetString("database.postgres.user"),
			Password: config.GetString("database.postgres.password"),
			Host:     config.GetString("database.postgres.host"),
			Port:     config.GetString("database.postgres.port"),
		},
		AWS: config.AWSConfig{
			AccessKey: config.GetString("aws.accessKey"),
			SecretKey: config.GetString("aws.secretKey"),
		},
		Redis: config.RedisConfig{
			Host: config.GetString("redis.host"),
			Port: config.GetString("redis.port"),
		},
	}
}
