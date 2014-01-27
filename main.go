package main

import(
	"fmt"
	//"net/http"
	//"github.com/nradz/DistGo/db"
	//"github.com/nradz/DistGo/rq"
	"github.com/nradz/DistGo/auxiliar"
)


func main() {

	fmt.Println("Starting DistGo...")



	//db.StartDB() //initialize the database

	aux := auxiliar.LoadConf()

	fmt.Println(aux["ip"])

	fmt.Println("DistGo is working!")



	//http.HandleFunc("/", rq.LoadAnt)
	//http.ListenAndServe(":8080", nil)

    	
}
