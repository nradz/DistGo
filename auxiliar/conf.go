package auxiliar

import(
	//"io/ioutil"
	//"encoding/json"
	//"log"
)


//Temporal default config

var ip string = "127.0.0.1"
var port string = "8080"
var cors string = "*"

//Load the config file into a string map
func LoadConf() map[string]string{

	var conf map[string]string

	conf = make(map[string]string)
	conf["ip"] = ip
	conf["port"] = port
	conf["cors"] = cors

	return conf


	// file, err := ioutil.ReadFile("conf/"+filename+".json")
	
	// if err != nil{
	// 	log.Fatal("The config file "+filename+".json doesn't exist.")
	// }
	
	// json.Unmarshal(file, conf)

}