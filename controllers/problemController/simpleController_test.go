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

	data := make([]string)
	data[0] = "6"

	err := NewResult(id, data)

	if err != nil{
		t.Error(err.Error())
	}


}

func TestUpdate(t *Testing.T){
	setup()

	id1 := 5
	id2 := 6

	NewRequest(5)
	NewRequest(6)

	data := make([]string)
	data[0] = "6"	

	err := NewResult(id1, data)

	alg, update, err := NewRequest(id2)


}

//aux funcs
func setup(){
	go SimpleProblemController(problems.Getproblem("pruebaProblem"))

}