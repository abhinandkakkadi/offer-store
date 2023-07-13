package main

import (
	"log"

	"github.com/abhinandkakkadi/offer-store/cmd/api/docs"
	config "github.com/abhinandkakkadi/offer-store/pkg/config"
	di "github.com/abhinandkakkadi/offer-store/pkg/di"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

func main() {

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "offer-store"
	docs.SwaggerInfo.Description = "OFFER COMPANY"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, userRepository, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start(userRepository)
	}

}
