package simple

import(
"testing"
"github.com/nradz/DistGo/problems"
"time"
	)

func TestClose(t *testing.T){
	sp := New(problems.GetProblem("pruebaProblem"))
	sp.Init()
	sp.Close()	
}

func TestFirstRequest(t *testing.T){
	sp := New(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id uint32 = 1

	var lastUpdate uint32 = 0

	alg, datos, nUpdate, err := sp.NewRequest(id, lastUpdate)

	if alg == ""{
		t.Error("No alg")
	}

	if datos.(int64) != 0{
		t.Error("No 0")
	}

	if nUpdate != 1{
		t.Error("NUpdate is not 1")
	}

	if err != nil{
		t.Error(err.Error())
	}
}

func TestNewResult(t *testing.T){
	sp := New(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id uint32 = 1

	var lastUpdate uint32 = 1

	datos := make([]string, 1)
	datos[0] = "6"

	err := sp.NewResult(id, datos, lastUpdate)

	if err != nil{
		t.Error(err.Error())
	}

}

func TestUpdate(t *testing.T){
	sp := New(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id1 uint32 = 5
	var lastUpdate1 uint32 = 1
	var id2 uint32 = 6
	var lastUpdate2 uint32 = 1

	datos := make([]string, 1)
	datos[0] = "6"	

	err := sp.NewResult(id1, datos, lastUpdate1)

	if err != nil{
		t.Error("No result:", err)
	}

	time.Sleep(100 *time.Millisecond)

	alg, update, nUpdate, _ := sp.NewRequest(id2, lastUpdate2)

	if alg != ""{
		t.Error("alg:", alg)
	}

	if nUpdate != 2{
		t.Error("Incorrect nUpdate:", nUpdate)
	}

	if 6 != update.(int64){
		t.Error("No update:", update)
	}


}

func TestStandby(t *testing.T){
	sp := New(problems.GetProblem("pruebaProblem"))
	sp.Init()
	defer sp.Close()

	var id1 uint32 = 5
	var lastUpdate1 uint32 = 1
	var id2 uint32 = 6
	var lastUpdate2 uint32 = 1

	datos := make([]string, 1)
	datos[0] = "6"

	aux := make(chan problems.Data)

	go func(){
		_, update, _, _ := sp.NewRequest(id1, lastUpdate1)
		aux <- update
	}()

	sp.NewResult(id2, datos, lastUpdate2)

	res := <- aux

	if 6 != res.(int64){
		t.Error("Incorrect update:", res)
	}
}
