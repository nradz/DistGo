package problems

import(
	"testing"
	//"fmt"
	"time"
)


func TestGetProblem(t *testing.T){

	res := GetProblem("pruebaProblem")

	if res == nil{
		t.Error(res)
	}
}


func TestStart(t *testing.T){

	prob := pruebaProblem{}

	c := make(chan ProblemUpdate)

	prob.Start(c)

}

func TestNewBetterResult(t *testing.T){

	prob := pruebaProblem{}

	c := make(chan ProblemUpdate)

	prob.Start(c)

	data := make([]string, 1)
	data[0] = "14"

	var lastUpdate uint32 = 1

	go prob.NewResult(data, lastUpdate)

	res := <- c //a problemUpdate

	if res.Data.(int64) != 14{
		t.Error(res.Data)

	}

}

func TestNewWorseResult(t *testing.T){
	prob := pruebaProblem{}

	c := make(chan ProblemUpdate)

	prob.Start(c)

	data := make([]string, 1)
	data[0] = "14"

	var lastUpdate uint32 = 1

	go prob.NewResult(data, lastUpdate)

	time.Sleep(100 * time.Millisecond)

	select{
	case <-c:
		break
	default:
		t.Error("empty channel")
		break
	}


	data[0] = "5"

	go prob.NewResult(data, lastUpdate)

	time.Sleep(100 * time.Millisecond)

	select{
	case <-c:
		t.Error("non-empty channel")
		break
	default:
		break
	}

	
}
