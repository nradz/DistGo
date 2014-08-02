package problems

import(
	"os"
	"decoder"
	)


func init(){
	AddProblem("perFlowACOProblem", &perFlowACOProblem{})
}

type perFlowACOProblem struct{

}

func (prob *perFlowACOProblem) Type() string{
	return "simple"
}

func (prob *perFlowACOProblem) Start(c chan ProblemUpdate) ProblemUpdate{

	return ProblemUpdate{}
}

func (prob *perFlowACOProblem) NewResult(data []string, lastUpdate uint32){
	return
}

func (prob *perFlowACOProblem) Loop(){
	return
}

//Private functions
func loadConf(){
	root := os.Getenv("HOME")
	file, err := os.Open(root+"/.DistGo/DistGo.conf")
	if err != nil{
		log.Fatal("perFlowACOProblem-> loadConf error:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
}