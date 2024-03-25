package server

import (
	"fmt"
	"log"
	"net/http"

	config "layout/_config"
	router "layout/internal/route"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func Server() {
	color.Green("Server starting tag...")

	// connect all resources
	resource := &config.Resource{}

	// // database
	// resource, err := config.CreateResource()
	// if err != nil {
	// 	time.Sleep(5 * time.Second)
	// 	return
	// }
	// defer resource.Close()

	// // rabbit mq
	// if err := resource.CreateMQ(); err != nil {
	// 	fmt.Println("rabbit_error:", err)
	// }

	// // redis
	// if err := resource.RedisConnect(); err != nil {
	// 	fmt.Println("redis :", err)
	// }
	// defer resource.CloseRedis()

	// route
	r := gin.New()
	r.Use(CORS)
	router.Router(r, resource)

	fmt.Println("Server listening on port 8080...")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	} else {
		c.AbortWithStatus(http.StatusOK)
		return
	}
}
