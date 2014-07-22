package problemController


import(
	"github.com/nradz/DistGo/problems"
	"github.com/nradz/DistGo/controllers/problemController/simple"
)

type data interface{}


//Problem controllers have to implement this interface.
type ProblemController interface{
	Init()
	NewRequest(uint32, uint32) (string, problems.Data, uint32, error)
	NewResult(uint32, []string, uint32) error
	Close()
}

//New return a new problemController.
//It will be right for the problem.
func New(prob problems.Problem) ProblemController{
	switch prob.Type(){
		case "simple":
			return simple.New(prob)
		default:
			return nil
	}
}