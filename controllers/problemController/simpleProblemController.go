package problemController

import(
	"fmt"
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
	Err Error
}

var(
	clientChan = make(chan *clientRequest)
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

		case update = <-problemChan:
			problemState.Update(update)

		case req = <-clientChan:
			if req.Data != nil{
				problemState.NewResult(req.Id, req.Data,
					problem problems.Problem, req.ResChan)
			} else{
				problemState.NewRequest(req.Id, req.ResChan)
			}
		
		}
	}

}


func NewRequest(id uint32){

	req := &problemControlRequest{id, nil, make(chan problemControlResponse)}
	
	clientChan <- req

	res := <- req.Response

	return res.Alg, res.Data, res.Err

}

func NewResult(id uint32, data []string){
	req := &problemControlRequest{id, data, make(chan problemControlResponse)}
	
	clientChan <- req

	res := <- req.Response

	return res.Err
}