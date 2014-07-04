package conf
import(
	//"io/ioutil"
	//"encoding/json"
	//"log"
)

//default configuration
var(
	ip string = "127.0.0.1"
	port string = "8080"
	cors string = "*"
	problem string = "pruebaProblem"

)


//Temporal useless func
func LoadConf(){
	return
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