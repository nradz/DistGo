package main

import(
	"fmt"
	//"github.com/nradz/DistGo/db"
	//"github.com/nradz/DistGo/distgo_types"
	"github.com/nradz/DistGo/conf"
	"github.com/nradz/DistGo/controllers/connectionController"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/problems"
	//"github.com/nradz/DistGo/channels"
)


func main() {


	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	//Start Controllers
	cli := clientController.NewClientController()
	cli.Init()

	problem := problems.GetProblem(conf.Problem())
	probCon := problemController.NewSimpleProblemController(problem)
	probCon.Init()

	con := connectionController.NewConnectionController(cli, probCon)
	
	fmt.Println("DistGo is working!")

	con.Init()
  	
}
