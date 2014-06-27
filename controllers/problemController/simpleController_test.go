package problemController

import(
"testing"
"github.com/nradz/DistGo/problems"
	)



func TestFirstRequest(t *testing.T){
	setup()

	var id uint32 = 5

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

func TestNewResult(t *testing.T){
	setup()

	var id uint32 = 5

	NewRequest(5)

	data := make([]string, 1)
	data[0] = "6"

	err := NewResult(id, data)

	if err != nil{
		t.Error(err.Error())
	}


}

func TestUpdate(t *testing.T){
	setup()

	var id1 uint32 = 5
	var id2 uint32 = 6

	NewRequest(5)
	NewRequest(6)

	data := make([]string, 1)
	data[0] = "6"	

	NewResult(id1, data)

	alg, update, _ := NewRequest(id2)

	if alg != ""{
		t.Error("alg")
	}

	if 6 != update.(uint32){
		t.Error("No update")
	}


}

//aux funcs
func setup(){
	go SimpleProblemController(problems.GetProblem("pruebaProblem"))

}