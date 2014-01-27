package auxiliar

import(
	//"io/ioutil"
	//"encoding/json"
	//"log"
)


//Temporal default config

var ip string = "127.0.0.1"

//Load a json configuration file named "filename" in the "conf" struct
func LoadConf() map[string]string{

	var conf map[string]string

	conf = make(map[string]string)
	conf["ip"] = ip

	return conf


	// file, err := ioutil.ReadFile("conf/"+filename+".json")
	
	// if err != nil{
	// 	log.Fatal("The config file "+filename+".json doesn't exist.")
	// }
	
	// json.Unmarshal(file, conf)

}