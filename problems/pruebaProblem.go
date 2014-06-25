package problems

import(
	"strconv"
	//"fmt"
	)

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

func (prob *pruebaProblem) Init(c chan ProblemUpdate) ProblemUpdate{

	prob.best = 0

	prob.c = c

	return ProblemUpdate{alg, prob.best}

}

func (prob *pruebaProblem) NewResult(data []string){

	n, _ := strconv.ParseInt(data[0], 0, 0)

	if n > prob.best{
		prob.best = n

		prob.c <- ProblemUpdate{"", n}
	}

}

func (prob *pruebaProblem) Loop(){
	return
}