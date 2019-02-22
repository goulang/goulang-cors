package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang-cors/server/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	router := gin.New()
	// router.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		fmt.Println(origin)
		config.AllowOrigins = []string{origin}
		return true
	}
	// config.AddAllowMethods("GET")

	router.Use(cors.New(config))
	router.Use(gin.Logger())


	
	// session
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Domain: "www.goulang.org",
		MaxAge: 7 * 24 * 3600,
	})
	router.Use(sessions.Sessions("goulang", store))



	loaderRouter(router)
	router.Run()
}
func loaderRouter(c *gin.Engine) {
	c.GET("topics", routes.GetTopics)
	c.POST("topics", routes.PostTopics)
	c.POST("login",routes.PostLogin)
}
