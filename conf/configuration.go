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

//default configuration
var(
	ip string
	port string
	cors string
	problem string
)


//Temporal useless func
func LoadConf(){
	root := os.Getenv("HOME")
	file, err := os.Open(root+"/.DistGo/conf.json")
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

func Ip() string{
	return ip
}

func Port() string{
	return port
}

func Cors() string{
	return cors
}

func Problem() string{
	return problem
}