package main

import (
	"fmt"

	"log"
	"net/http"

	routes "github.com/rodrigo-brito/bus-api-go/action/cli/server/http"
	_ "github.com/rodrigo-brito/bus-api-go/config"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	defer mysql.CloseConnection()
	router := routes.InjectAPIRoutes()

	handler := cors.Default().Handler(router)

	fmt.Printf("Server started at http://localhost:%d\n", viper.GetInt("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), handler))
}
