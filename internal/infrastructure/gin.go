package infrastructure

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/MSHR-Dec/go_backend/internal/infrastructure/middleware"
)

func InitGin(db *gorm.DB, cache redis.Store) *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("patune_token", cache))

	setLogger(r)
	setCORS(r)
	setRoutes(r, db)

	return r
}

func setLogger(r *gin.Engine) {
	logger, _ := zap.NewProduction()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}

func setCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:63342",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
	}))
}

func setRoutes(r *gin.Engine, db *gorm.DB) {
	userController := initializeUser(db)
	user := r.Group("/user")
	{
		user.Use(middleware.SetSession())
		user.POST("/signin", func(context *gin.Context) {
			userController.SignIn(context)
		})
	}

	taskController := initializeTask(db)
	task := r.Group("/tasks")
	{
		task.Use(middleware.CheckLoggedIn())
		task.GET("/:userID", func(context *gin.Context) {
			taskController.ListTasksByUserID(context)
		})
	}
}
