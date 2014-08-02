package problems

import(
		"github.com/nradz/DistGo/conf"
		"time"
		"fmt"
		"os"
		"log"
		)

//NewLog return a new logger file. The Path is 
//$HOME/.DistGo/'ProblemName'/'Date'/'name'+'hour'
//For example: $HOME/.DistGo/pruebaProblem/01-02-2014/valores-14-45-10
func NewLog(name string, prefix string, flag int) *log.Logger{
	root := os.Getenv("HOME")
	t := time.Now()

	date := t.Format("1-2-2006")

	hour := t.Format("-15-04-05")

	filename := name+hour

	path := fmt.Sprintf("/%s/.DistGo/%s/%s/", root, conf.Problem(), date)

	//Check if the directory exists. Otherwise, create it.
	_, err := os.Stat(path)
	if err != nil{
		if os.IsNotExist(err){
			os.Mkdir(path, os.ModePerm)
		} else{
			log.Fatal("Error en NewLog-1: ", err)
		}
	}

	f, err := os.Create(path+filename)
	if err != nil{
		log.Fatal("Error en NewLog-2: ", err)
	}

	return log.New(f, prefix, flag)
}