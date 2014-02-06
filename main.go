package main

import(
	"fmt"
	"net/http"
	//"github.com/nradz/DistGo/db"
	//"github.com/nradz/DistGo/distgo_types"
	"github.com/nradz/DistGo/auxiliar"
	"github.com/nradz/DistGo/controllers/connectionController"
)


func main() {


	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	conf := auxiliar.LoadConf()

	//Start 

	fmt.Println("DistGo is working!")


	http.HandleFunc("/",connectionController.Controller(conf))
	http.ListenAndServe(":"+conf["port"], nil)

    	
}
