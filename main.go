package main

import(
	"fmt"
	//"github.com/nradz/DistGo/db"
	//"github.com/nradz/DistGo/distgo_types"
	"github.com/nradz/DistGo/configuration"
	"github.com/nradz/DistGo/controllers/connectionController"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/problems"
	//"github.com/nradz/DistGo/channels"
)


func main() {


	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	//Load conf
	configuration.LoadConf()

	//Start Controllers
	go clientController.ClientController()

	problem := problems.GetProblem(configuration.Configuration().Problem())

	go problemController.SimpleProblemController(problem)

	fmt.Println("DistGo is working!")

	connectionController.ConnectionController()
  	
}
