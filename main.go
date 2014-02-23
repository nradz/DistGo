package main

import(
	"fmt"
	"net/http"
	//"github.com/nradz/DistGo/db"
	//"github.com/nradz/DistGo/distgo_types"
	"github.com/nradz/DistGo/configuration"
	"github.com/nradz/DistGo/controllers/connectionController"
	"github.com/nradz/DistGo/controllers/clientController"
	//"github.com/nradz/DistGo/channels"
)


func main() {


	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	//Load conf
	configuration.LoadConf()

	//Start Controllers
	go clientController.ClientController()

	fmt.Println("DistGo is working!")

	http.HandleFunc("/",connectionController.ConnectionController())
	http.ListenAndServe(":8080", nil)

    	
}
