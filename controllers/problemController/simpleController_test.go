package problemController

import(
"testing"
"problems"
	)



func TestFirstRequest(t *Testing.T){
	setup()

	id := 5

	alg, data, err := NewRequest(id)

	if alg == ""{
		t.Error("No alg")
	}

	if data.(int64) != 0{
		t.Error("No 0")
	}

	if err != nil{
		t.Error(err.Error())
	}
}

func TestNewResult(t *Testing.T){
	setup()

	id := 5

	NewRequest(5)

	err := NewResult(id, data)

	if err != nil{
		t.Error(err.Error())
	}


}

func TestFirstRequest(t *Testing.T){}

func TestFirstRequest(t *Testing.T){}

func TestFirstRequest(t *Testing.T){}

func TestFirstRequest(t *Testing.T){}

func TestFirstRequest(t *Testing.T){}

func TestFirstRequest(t *Testing.T){}

func TestFirstRequest(t *Testing.T){}


//aux funcs
func setup(){
	go SimpleProblemController(problems.Getproblem("pruebaProblem"))

}