package db


import(
	"fmt"
	"github.com/nradz/DistGo/utils"
)


type Conf struct{
	Name string
	Data []string

}

func StartDB(){
	
	conf := &Conf{}
	utils.LoadConf("db", conf)

	fmt.Println(conf.Name)
}