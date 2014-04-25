package problems

import(
	"fmt"
	"github.com/nradz/DistGo/channels"
	)

type pruebaProblem struct{

}

const(
	alg = `function mainFunc(romero,data){
				window.alert("miau");
					romero.finish();
				}
	`	
	)

func (prob pruebaProblem) Init() channels.ProblemUpdate{

	return channels.ProblemUpdate{alg, nil}

}

func (prob pruebaProblem) NewResult(data []string){

	fmt.Println(data)

}

func (prob pruebaProblem) Loop(){
	return
}