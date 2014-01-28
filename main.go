package main

import(
	"fmt"
	//"net/http"
	//"github.com/nradz/DistGo/db"
	"github.com/nradz/DistGo/distgo_types"
	//"github.com/nradz/DistGo/auxiliar"
)


func main() {

	fmt.Println("Starting DistGo...")

	//db.StartDB() //initialize the database

	//conf := auxiliar.LoadConf()

	clientList := distgo_types.NewClientList()

	clientList.NewClient("127.0.0.1")

	nuevoCliente(clientList)

	for k, v := range clientList.List{
		fmt.Println(k)
		fmt.Println(v)
	}

	fmt.Println("DistGo is working!")


	//http.HandleFunc("/", rq.LoadProblem(conf))
	//http.ListenAndServe(":"+conf["port"], nil)

    	
}


func nuevoCliente(a distgo_types.ClientList){

	a.NewClient("127.0.0.1")
}
