package main

import(
	"fmt"
	"net/http"
	//"github.com/nradz/DistGo/db"
	"github.com/nradz/DistGo/rq"
	"github.com/nradz/DistGo/auxiliar"
)


func main() {

	fmt.Println("Starting DistGo...")



	//db.StartDB() //initialize the database

	conf := auxiliar.LoadConf()

	fmt.Println("DistGo is working!")


	http.HandleFunc("/", rq.LoadProblem(conf))
	http.ListenAndServe(":"+conf["port"], nil)

    	
}
