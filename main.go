package main

import(
	"fmt"
	"log"
	"errors"
	"runtime"
	"github.com/nradz/DistGo/conf"
	"github.com/nradz/DistGo/controllers/connectionController"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/problems"
	"github.com/nradz/DistGo/problems/pruebaProblem"
	"github.com/nradz/DistGo/problems/perFlowACOProblem"
)

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {


	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	//Start Controllers
	cli := clientController.New()
	cli.Init()

	problem, err := getProblem(conf.Problem())	
	if err != nil{
		log.Fatal(err)
	}

	probCon, err := problemController.New(problem)
	if err != nil{
		log.Fatal(err)
	}

	probCon.Init()

	con := connectionController.New(cli, probCon)
	
	fmt.Println("DistGo is working!")

	con.Init()
  	
}


//Return a problem by his name
func getProblem(prob string) (problems.Problem, error){
	switch prob{
	case "pruebaProblem":
		return pruebaProblem.New(), nil
	case "perFlowACOProblem":
		return perFlowACOProblem.New(), nil
	default:
		return nil, errors.New("Problem " + prob + " not found.")
	}
}