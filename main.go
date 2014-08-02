package main

import(
	"fmt"
	"github.com/nradz/DistGo/conf"
	"github.com/nradz/DistGo/controllers/connectionController"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/problems"
)


func main() {


	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	//Start Controllers
	cli := clientController.New()
	cli.Init()

	problem := problems.GetProblem(conf.Problem())
	probCon := problemController.New(problem)
	probCon.Init()

	con := connectionController.New(cli, probCon)
	
	fmt.Println("DistGo is working!")

	con.Init()
  	
}
