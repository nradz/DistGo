package problems

import(
	"testing"
	//"fmt"
	"time"
)

func TestInit(t *testing.T){

	prob := pruebaProblem{}

	c := make(chan ProblemUpdate)

	prob.Init(c)

}

func TestNewBetterResult(t *testing.T){

	prob := pruebaProblem{}

	c := make(chan ProblemUpdate)

	prob.Init(c)

	data := make([]string, 1)
	data[0] = "14"

	go prob.NewResult(data)

	res := <- c

	if res.Data.(int64) != 14{
		t.Error(res.Data)

	}

}

func TestNewWorseResult(t *testing.T){
	prob := pruebaProblem{}

	c := make(chan ProblemUpdate)

	prob.Init(c)

	data := make([]string, 1)
	data[0] = "14"

	go prob.NewResult(data)

	time.Sleep(100 * time.Millisecond)

	select{
	case <-c:
		break
	default:
		t.Error("empty channel")
		break
	}


	data[0] = "5"

	go prob.NewResult(data)

	time.Sleep(100 * time.Millisecond)

	select{
	case <-c:
		t.Error("non-empty channel")
		break
	default:
		break
	}

	
}
