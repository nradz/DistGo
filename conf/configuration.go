//Package conf loads automatically the configuration
//of DistGo from the file "$HOME/.DistGo/DistGo.conf"
package conf

import(
	"os"
	"encoding/json"
	"log"
)


type configFile struct{
	NClients uint16
	Ip string
	Port string
	Cors string
	Problem string
}

//configuration vars
var(
	nClients uint16
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

	nClients = data.NClients
	ip = data.Ip
	port = data.Port
	cors = data.Cors
	problem = data.Problem
}

func NClients() uint16{
	return nClients
}

//Return the ip of the server
func Ip() string{
	return ip
}

//Return the port where DistGo is listening
func Port() string{
	return port
}

//Return the content of the header field
//"Access-Control-Allow-Origin" of the responses
func Cors() string{
	return cors
}

//Return the name of the problem that is being executed
func Problem() string{
	return problem
}