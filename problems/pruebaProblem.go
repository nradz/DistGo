package problems

import(
	"fmt"
	"github.com/nradz/DistGo/channels"
	)

type pruebaProblem struct{

}

const(
	alg = `window.alert("miau");
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