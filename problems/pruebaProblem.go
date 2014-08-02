package problems

import(
	"strconv"
	"os"
	"io/ioutil"
	"log"
	//"fmt"
	)


func init(){
	AddProblem("pruebaProblem", &pruebaProblem{})
}

type pruebaProblem struct{
	alg string
	best int64
	number uint32
	file *log.Logger
	c chan ProblemUpdate
}

func (prob *pruebaProblem) Type() string{
	return "simple"
}

func (prob *pruebaProblem) Start(c chan ProblemUpdate) ProblemUpdate{

	prob.file = NewLog("valores", "", 2)

	prob.best = 0

	prob.c = c

	prob.number = 1

	//Load the algorithm
	root := os.Getenv("HOME")
	buf, err := ioutil.ReadFile(root+"/.DistGo/pruebaProblem/alg.js")
	if err != nil{
		log.Fatal("prueba Start error: ", err)
	}

	prob.alg = string(buf)



	return ProblemUpdate{prob.alg, prob.best, prob.number}

}

func (prob *pruebaProblem) NewResult(data []string, lastUpdate uint32){

	n, _ := strconv.ParseInt(data[0], 0, 0)

	if n > prob.best{
		prob.file.Println(n)
		prob.best = n
		prob.number += 1

		prob.c <- ProblemUpdate{"", n, prob.number}
	}

}

func (prob *pruebaProblem) Loop(){
	return
}