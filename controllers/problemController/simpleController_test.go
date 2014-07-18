package problemController

import(
"testing"
"github.com/nradz/DistGo/problems"
"time"
	)

func TestClose(t *testing.T){
	sp := NewSimpleProblemController(problems.GetProblem("pruebaProblem"))
	sp.Init()
	sp.Close()	
}

func TestFirstRequest(t *testing.T){
	sp := NewSimpleProblemController(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id uint32 = 1

	var lastUpdate uint32 = 0

	alg, datos, NUpdate, err := sp.NewRequest(id, lastUpdate)

	if alg == ""{
		t.Error("No alg")
	}

	if datos.(int64) != 0{
		t.Error("No 0")
	}

	if NUpdate != 1{
		t.Error("NUpdate is not 1")
	}

	if err != nil{
		t.Error(err.Error())
	}
}

func TestNewResult(t *testing.T){
	sp := NewSimpleProblemController(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id uint32 = 1

	var lastUpdate uint32 = 0

	sp.NewRequest(id, lastUpdate)

	datos := make([]string, 1)
	datos[0] = "6"

	err := sp.NewResult(id, lastUpdate, datos)

	if err != nil{
		t.Error(err.Error())
	}

}

func TestUpdate(t *testing.T){
	sp := NewSimpleProblemController(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id1 uint16 = 5
	var id2 uint16 = 6

	sp.NewRequest(id1)
	sp.NewRequest(id2)

	datos := make([]string, 1)
	datos[0] = "6"	

	err := sp.NewResult(id1, datos)

	if err != nil{
		t.Error("No result:", err)
	}


	time.Sleep(100 *time.Millisecond)

	alg, update, _ := sp.NewRequest(id2)

	if alg != ""{
		t.Error("alg")
	}

	if 6 != update.(int64){
		t.Error("No update:", update)
	}


}
