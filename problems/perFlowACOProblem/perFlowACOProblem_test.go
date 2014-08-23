package perFlowACOProblem

import(
	//"fmt"
	"reflect"
	"sort"
	"testing"
	"github.com/nradz/DistGo/conf"
	//"github.com/nradz/DistGo/problems"
)

var costs_test [][]int

func setup(){
	conf.SetProblem("perFlowACOproblem")
	costs_test = [][]int{{54,83,15,71,77},
	{79,3,11,99,56},
	{16,89,49,15,89},
	{66,58,31,68,78},
	{58,56,20,85,53}}
}

func TestLoadConf(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.loadConf()
}

func TestLoadProblem(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.loadConf()
	prob.loadProblem()
}

func TestEvaluate(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.jobs = 20
	prob.machines = 5
	prob.costs = costs_test

	total := prob.evaluate([]int{4,3,2})
	
	if total != 1279 {
		t.Error("Incorrect total: ", total)
	}
}

func TestVerifySeq(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.jobs = 5
	prob.machines = 5
	prob.costs = costs_test

	ok := prob.verifySeq([]int{0,4,2,3,1})
	if !ok{
		t.Error("No ok")
	}

	ok = prob.verifySeq([]int{0,4,3,3,1})
	if ok{
		t.Error("ok")
	}
}

func TestByCosts(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.jobs = 5
	prob.machines = 5
	prob.costs = costs_test

	seq := []int{0,1,2,3,4}
	sort.Sort(sort.Reverse(prob.byCosts(seq)))
	

	ok := reflect.DeepEqual(seq, []int{4,3,1,0,2})

	if !ok{
		t.Error("Incorrect sequence:", seq)
	}

}

func TestByTotalFlowtime(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.jobs = 5
	prob.machines = 5
	prob.costs = costs_test

	list := [][]int{{1,2},{2,1}}

	sort.Sort(prob.byTotalFlowtime(list))

	ok := reflect.DeepEqual(list, [][]int{{2,1},{1,2}})

	if !ok{
		t.Error("Incorrect order:", list)
	}
}

func TestNEH(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.jobs = 5
	prob.machines = 5
	prob.costs = costs_test

	seq := prob.neh()

	ok := reflect.DeepEqual(seq, []int{2,0,1,3,4})

	if !ok{
		t.Error("Incorrect order:", seq)
	}
	
}

func TestJobIndexBased(t *testing.T){
	setup()
	prob := perFlowACOProblem{}
	prob.jobs = 5
	prob.machines = 5
	prob.costs = costs_test

	res := prob.jobIndexBased([]int{0,1,2,3,4})
	
	ok := reflect.DeepEqual(res, []int{2,0,1,3,4})

	if !ok{
		t.Error("Incorrect order: ", res)
	}
}

// func TestStart(t *testing.T){
// 	setup()
// 	prob := perFlowACOProblem{}

// 	c := make(chan problems.ProblemUpdate)

// 	prob.Start(c)

// 	//res := prob.Start(c)

// 	//fmt.Println(res)
// }

// func TestNewResult(t *testing.T){
// 	setup()
// 	prob := perFlowACOProblem{}
// 	prob.Start(make(chan problems.ProblemUpdate))

// 	prob.jobs = 5
// 	prob.machines = 5
// 	prob.costs = costs_test
// 	prob.bestSeq = []int{0,1,2,3,4}
// 	prob.status = 1
// 	prob.checkTime = 3000
// 	prob.newSeq = make(chan []int)
// 	prob.updateChan = make(chan problems.ProblemUpdate)

// 	go prob.Loop()

// 	go prob.NewResult([]string{"2","0","1","3","4"}, 2)

// 	res := <- prob.updateChan
// 	fmt.Println(res)


// }

// func TestIncorrectResult(t *testing.T){

// 	setup()
// 	prob := perFlowACOProblem{}
// 	prob.Start(make(chan problems.ProblemUpdate))

// 	prob.jobs = 5
// 	prob.machines = 5
// 	prob.costs = costs_test
// 	prob.bestSeq = []int{0,1,2,3,4}
// 	prob.status = 1
// 	prob.checkTime =3000
// 	prob.newSeq = make(chan []int)
// 	prob.updateChan = make(chan problems.ProblemUpdate)

// 	go prob.NewResult([]string{"0","1","1","3","4"}, 2)
// }