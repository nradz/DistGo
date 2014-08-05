package perFlowACOProblem

import(
	//"fmt"
	"time"
	"strconv"
	"reflect"
	"sort"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/nradz/DistGo/problems"
	)

type configFile struct{
	Machines int
	Jobs int
	Number int
	CheckTime int
}

type taillard struct{
	Taillard []struct{
		Machines int
		Jobs int
		Problems [][][]int
	}
}


type perFlowACOProblem struct{
	machines int //Set by loadConf
	jobs int //Set by loadConf
	number int //Set by loadConf
	costs [][]int //Set by loadProblem
	bestSeq []int //The best sequence received
	receivedSeqs [][]int //List of received sequences 
	//that will be evaluated
	status uint32 //The status number of the problem
	alg string //The algorithm that will be 
	//executed by the clients
	updateChan chan problems.ProblemUpdate //The chan 
	//where problem sends updates to the controller
	newSeq chan []int //The chan where NewResult sends 
	//data to the Loop
	checkTime int //The time between evaluations of 
	//received sequences
	timer *time.Timer //Timer that indicate when 
	//evaluated the received sequences

}


func init(){
	problems.AddProblem("perFlowACOProblem", &perFlowACOProblem{})
}

func (prob *perFlowACOProblem) Type() string{
	return "simple"
}

func (prob *perFlowACOProblem) Start(c chan problems.ProblemUpdate) problems.ProblemUpdate{
	prob.loadConf()
	prob.loadProblem()

	prob.bestSeq = prob.neh()
	prob.status = 1
	prob.newSeq = make(chan []int)
	prob.updateChan = c

	//Load the algorithm
	root := os.Getenv("HOME")
	buf, err := ioutil.ReadFile(root+"/.DistGo/perFlowACOProblem/alg.js")
	if err != nil{
		log.Fatal("perFlowACO Start error: ", err)
	}

	prob.alg = "var costs = " + prob.matrix2string(prob.costs) + "\n" + string(buf)

	return problems.ProblemUpdate{prob.alg, prob.bestSeq, prob.status}
}

func (prob *perFlowACOProblem) NewResult(data []string, lastUpdate uint32){
	//Transform []string to []int
	seq := make([]int, prob.jobs)
	for i, v := range data{
		conv, err := strconv.Atoi(v)
		if err != nil{
			return
		}
		seq[i] = conv
	}

	//Verify seq
	if prob.verifySeq(seq){
		prob.newSeq <- seq
	}
}

func (prob *perFlowACOProblem) Loop(){
	prob.timer = time.NewTimer(time.Duration(prob.checkTime) * time.Millisecond)

	for{
		select{
		case seq := <-prob.newSeq:

			prob.receivedSeqs = append(prob.receivedSeqs, seq)

		case <- prob.timer.C:

			if len(prob.receivedSeqs) > 0{
				prob.receivedSeqs = append(prob.receivedSeqs, prob.bestSeq)
				sort.Sort(prob.byTotalFlowtime(prob.receivedSeqs))

				//If the best seq of the list is NOT the current best seq
				if !reflect.DeepEqual(prob.receivedSeqs[0], prob.bestSeq){
					prob.bestSeq = prob.receivedSeqs[0]
					prob.status += 1
					prob.updateChan <- problems.ProblemUpdate{"", prob.bestSeq, prob.status}
				}

				//Reset receivedSeq
				prob.receivedSeqs = nil
			}

			prob.timer.Reset(time.Duration(prob.checkTime) * time.Millisecond)

		}
	}
}



//Private functions
func (prob *perFlowACOProblem) loadConf(){
	root := os.Getenv("HOME")
	file, err := os.Open(root+"/.DistGo/perFlowACOProblem/conf.json")
	if err != nil{
		log.Fatal("perFlowACOProblem-> loadConf error:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	data := configFile{}

	err = decoder.Decode(&data)
	if err != nil{
		log.Fatal("perFlowACOProblem-> loadConf error:", err)
	}

	prob.machines = data.Machines
	prob.jobs = data.Jobs
	prob.number = data.Number
	prob.checkTime = data.CheckTime

}

func (prob *perFlowACOProblem) loadProblem(){
	root := os.Getenv("HOME")
	file, err := os.Open(root+"/.DistGo/perFlowACOProblem/taillard.json")
	if err != nil{
		log.Fatal("perFlowACOProblem-> loadProblem error:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	data := taillard{}

	err = decoder.Decode(&data)
	if err != nil{
		log.Fatal("perFlowACOProblem-> loadProblem error:", err)
	}

	matrix := [][]int{}

	for _, v := range data.Taillard{
		if v.Machines == prob.machines && v.Jobs == prob.jobs{
			matrix = v.Problems[prob.number]
			break
		}
	}

	prob.costs = matrix
}

func (prob *perFlowACOProblem) evaluate(sequence []int) int{
	
	flowSeq := make([]int, len(sequence))

	for m := 0; m < prob.machines; m++{

		for i, v := range sequence{
			switch{
			
			case i == 0 && m == 0:
			flowSeq[0] = prob.costs[0][v]
			
			case i == 0:
			flowSeq[i] = flowSeq[i] + prob.costs[m][v]

			default:
			flowSeq[i] = Math.max(flowSeq[i-1], flowSeq[i]) + prob.costs[m][v]
			}
		}

	}

	total := 0
	for _, v := range flowSeq{
		total += v
	}

	return total
}

func (prob *perFlowACOProblem) neh() []int{
	//Make the sequence
	seq := make([]int, prob.jobs)
	for i := 0; i < prob.jobs; i++{
		seq[i] = i
	}

	//Order the sequence
	sort.Sort(sort.Reverse(prob.byCosts(seq)))

	//Step 2
	aux := [][]int{{seq[0], seq[1]},{seq[1], seq[0]}}

	sort.Sort(prob.byTotalFlowtime(aux))

	res := aux[0]

	if prob.jobs == 2{
		return res
	}

	for i := 2; i <prob.jobs; i++{
		list := [][]int{}

		for e := 0; e < (i+1); e++{
			aux := make([]int, len(res))
			copy(aux,res)
			switch{				
			case e == 0:
				aux = append([]int{seq[i]}, res...)
			case e == i:
				aux = append(aux, seq[i])
			default:
				aux = append(aux[:e], seq[i])
				aux = append(aux, res[e:]...)
			}
			list = append(list, aux)

		}

		sort.Sort(prob.byTotalFlowtime(list))
		res = list[0]
	}	

	return res
}

func (prob *perFlowACOProblem) verifySeq(seq []int) bool{
	
	//Check length
	if len(seq) != prob.jobs{
		return false
	}

	//Check if all the jobs are in the sequence
	for i := 0; i < prob.jobs; i++{
		
		inSeq := false
		
		for _, v := range seq{
			if i == v{
				inSeq = true
				break
			}
		}
		
		if !inSeq{
			return false
		} else{
			inSeq = false
		}
		
	}

	//If all the jobs are in the sequence
	return true
}

func (prob *perFlowACOProblem) byCosts(seq []int) *byCosts{
	res := byCosts{}
	res.seq = seq
	res.costs = make([]int, len(seq))

	for i, v := range seq{
		acum := 0
		for m:= 0; m < prob.machines; m++{
			acum += prob.costs[m][v]
		}
		res.costs[i] = acum
	}

	return &res
}

func (prob *perFlowACOProblem) byTotalFlowtime(list [][]int) *byTotalFlowtime{
	res := byTotalFlowtime{}
	res.list = list
	res.flowtimes = make([]int, len(list))

	for i, _ := range list{
		res.flowtimes[i] = prob.evaluate(list[i])
	}
	
	return &res	
}

func (prob *perFlowACOProblem) matrix2string(list [][]int) string{
	res := "["

	for i, v := range list{
		res += "["

		for e, num := range v{

			res += strconv.Itoa(num)
			if e != (len(v)-1){
				res += ","
			}
		}

		res += "]"
		if i != (len(list)-1){
			res += ",\n"
		}
		
	}

	res += "];\n" 

	return res
}

func max(x,y int) int{
	if x > y{
		return x
	} else{
		return y
	}
}
