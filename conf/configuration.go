//Package conf loads automatically the configuration
//of DistGo from the file "$HOME/.DistGo/DistGo.conf"
package conf

import(
	"os"
	"encoding/json"
	"log"
)


type configFile struct{
	MaxClients uint32
	Ip string
	Port string
	Cors string
	Problem string
}

//configuration vars
var(
	maxClients uint32
	ip string
	port string
	cors string
	problem string
)


//Load the configuration from the file 
//"$HOME/.DistGo/DistGo.conf"
func init(){
	root := os.Getenv("HOME")
	file, err := os.Open(root+"/.DistGo/DistGo.conf")
	if err != nil{
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	data := configFile{}

	err = decoder.Decode(&data)
	if err != nil{
		log.Fatal(err)
	}

	maxClients = data.MaxClients
	ip = data.Ip
	port = data.Port
	cors = data.Cors
	problem = data.Problem
}

//NClients returns the maximum number of clients
func MaxClients() uint32{
	return maxClients
}

func SetMaxClients(num uint32){
	maxClients = num
} 

//Ip returns the ip of the server
func Ip() string{
	return ip
}

//Port returns the port where DistGo is listening
func Port() string{
	return port
}

//Cors returns the content of the header field
//"Access-Control-Allow-Origin" of the responses
func Cors() string{
	return cors
}

//Problem returns the name of the problem that is being executed
func Problem() string{
	return problem
}