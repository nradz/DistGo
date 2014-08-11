package pruebaProblem

import(
	"github.com/nradz/DistGo/problems"
	"github.com/nradz/DistGo/conf"
	"testing"
	//"fmt"
	"time"
)

func init(){
	conf.SetProblem("pruebaProblem")
}

func TestGetProblem(t *testing.T){

	res := problems.GetProblem("pruebaProblem")

	if res == nil{
		t.Error(res)
	}

}

func TestType(t *testing.T){

	prob := pruebaProblem{}

	res := prob.Type()

	if res != "simple"{
		t.Error("Not match:", res)
	}
}

func TestStart(t *testing.T){

	prob := pruebaProblem{}

	c := make(chan problems.ProblemUpdate)

	prob.Start(c)

}

func TestNewBetterResult(t *testing.T){

	prob := pruebaProblem{}

	c := make(chan problems.ProblemUpdate)

	prob.Start(c)

	data := make([]string, 1)
	data[0] = "14"

	var lastUpdate uint32 = 1

	go prob.NewResult(data, lastUpdate)

	res := <- c //a problemUpdate

	if res.Number != 2{
		t.Error("Incorrect Number:", res.Number)
	}
	if res.Data.(int64) != 14{
		t.Error(res.Data)

	}

}

func TestNewWorseResult(t *testing.T){
	prob := pruebaProblem{}

	c := make(chan problems.ProblemUpdate)

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
