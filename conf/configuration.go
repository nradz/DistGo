//Package conf loads automatically the configuration
//of DistGo from the file "$HOME/.DistGo/DistGo.conf"
package conf

import(
	"os"
	"encoding/json"
	"log"
)


type configFile struct{
	Ip string
	Port string
	Cors string
	Problem string
}

//configuration vars
var(
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

	ip = data.Ip
	port = data.Port
	cors = data.Cors
	problem = data.Problem
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