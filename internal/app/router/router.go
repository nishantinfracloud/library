package router

import (
	"library/internal/app/constants"
	"strings"

	cors "github.com/rs/cors/wrapper/gin"
	"github.com/spf13/viper"

	controller "library/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

//NewRouter :
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	trustedProxies := viper.GetString(constants.TrustedProxies)
	trustedProxiesList := strings.Split(trustedProxies, ",")
	router.SetTrustedProxies(trustedProxiesList)

	allowedOrigins := viper.GetString(constants.AllowedOrigins)
	allowedOriginsList := strings.Split(allowedOrigins, ",")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: allowedOriginsList,
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"},
	})

	router.Use(corsHandler)

	library := router.Group(constants.Library)
	{
		api := library.Group(constants.API)
		{
			v1 := api.Group(constants.Version_V1)
			{
				v1.GET(constants.FetchBooks, controller.FetchBooks())
				v1.POST(constants.AddBook, controller.AddBook())
			}
		}
	}
	return router
}
