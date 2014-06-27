package problemController

import(
	//"fmt"
	"github.com/nradz/DistGo/configuration"
	"github.com/nradz/DistGo/problems"
	)

type data interface{}

type problemControlRequest struct{
	Id uint32
	Data []string
	Response chan problemControlResponse
}

type problemControlResponse struct{
	Alg string
	Data data
	Err error
}

var(
	clientChan = make(chan *problemControlRequest)
	problemChan = make(chan problems.ProblemUpdate)
	conf = configuration.Configuration()
	)

var problemState = simpleProblemState{}


func SimpleProblemController(problem problems.Problem){

	problemState = newSimpleProblemState()

	firstUpdate := problem.Init(problemChan)

	problemState.Alg = firstUpdate.Alg
	problemState.LastUpdate = firstUpdate.Data

	go problem.Loop()//Asynchronous execution of the problem loop (if it exists)


	var update = problems.ProblemUpdate{}
	var req = &problemControlRequest{}

	for{
		select{

		case update = <- problemChan:
			problemState.Update(update)

		case req = <- clientChan:
			if req.Data != nil{
				problemState.NewResult(req.Id, req.Data,
				 problem, req.Response)
			} else{
				problemState.NewRequest(req.Id, req.Response)
			}
		
		}
	}

}


func NewRequest(id uint32) (string, data, error){

	req := &problemControlRequest{id, nil, make(chan problemControlResponse)}
	
	clientChan <- req

	res := <- req.Response

	return res.Alg, res.Data, res.Err

}

func NewResult(id uint32, data []string) error{
	req := &problemControlRequest{id, data, make(chan problemControlResponse)}
	
	clientChan <- req

	res := <- req.Response

	return res.Err
}