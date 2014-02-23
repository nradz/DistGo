package configuration
import(
	//"io/ioutil"
	//"encoding/json"
	//"log"
)

var conf = configuration{}

type configuration struct{
	ip string
	port string
	cors string
}


//Temporal default data
func LoadConf(){
	conf.ip = "127.0.0.1"
	conf.port = "8080"
	conf.cors = "*"
}

func Configuration() configuration{
	return conf
}

func (c *configuration) Ip() string{
	return conf.ip
}

func (c *configuration) Port() string{
	return conf.port
}

func (c *configuration) Cors() string{
	return conf.cors
}