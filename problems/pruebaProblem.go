package problems

import(
	"strconv"
	//"fmt"
	)


func init(){
	AddProblem("pruebaProblem", &pruebaProblem{})
}

type pruebaProblem struct{
	best int64
	c chan ProblemUpdate
}

const(
	alg = `function mainFunc(romero,data){
				window.alert("miau");
					romero.finish();
				}
	`	
	)

func (prob *pruebaProblem) Start(c chan ProblemUpdate) ProblemUpdate{

	prob.best = 0

	prob.c = c

	return ProblemUpdate{alg, prob.best}

}

func (prob *pruebaProblem) NewResult(data []string, lastUpdate uint32){

	n, _ := strconv.ParseInt(data[0], 0, 0)

	if n > prob.best{
		prob.best = n

		prob.c <- ProblemUpdate{"", n}
	}

}

func (prob *pruebaProblem) Loop(){
	return
}