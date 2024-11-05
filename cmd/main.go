package main

import (
	"os"
	"time"

	"github.com/elonghi/encurtago/internal/controller"
	"github.com/elonghi/encurtago/internal/infrastructure/caching"
	"github.com/elonghi/encurtago/internal/infrastructure/database"
	"github.com/elonghi/encurtago/internal/repository"
	"github.com/elonghi/encurtago/internal/routes"
	"github.com/elonghi/encurtago/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	caching.StartRedis()
	godotenv.Load()

	linkRepo := repository.NewLinkRepository(database.DB, caching.RedisClient)
	userService := service.NewLinkService(linkRepo)
	linkController := controller.NewLinkController(userService)

	router := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	routes.InitializeRoutes(router, linkController)

	router.Run(os.Getenv("PORT"))
}
