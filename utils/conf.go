package utils

import(
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
)


type Conf interface{}

func LoadConf(filename string, conf Conf){
	file, err := ioutil.ReadFile("conf/"+filename+".json")
	
	if err != nil{
		log.Fatal("The config file "+filename+".json doesn't exist.")
	}
	
	json.Unmarshal(file, conf)
	
}