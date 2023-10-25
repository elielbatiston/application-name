package main

import (
	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
	"github.com/elielbatiston/application-name/src/application/controllers"
)

func init() {
	colibri.InitializeApp()
	//storage.Initialize() // uncomment if you use storage
	//cacheDB.Initialize() // uncomment if you use cache
	//sqlDB.Initialize() // uncomment if you use sql database
	//messaging.Initialize() // uncomment if you use messaging
	// application-name
}

func main() {
	restserver.AddRoutes(controllers.NewDemoController().Routes())
	restserver.ListenAndServe()
}
