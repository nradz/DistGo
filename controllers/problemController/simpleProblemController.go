package problemController

import(
	//"fmt"
	"errors"
	//"github.com/nradz/DistGo/conf"
	"github.com/nradz/DistGo/problems"
	)

type SimpleProblemController struct{
	clientChan chan *problemControlRequest
	problemChan chan problems.ProblemUpdate
	closeChan chan bool
	problem problems.Problem
	problemState *simpleProblemState
}

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

func NewSimpleProblemController(prob problems.Problem) *SimpleProblemController{
	return &SimpleProblemController{
		problem: prob,
	}

}


func (s *SimpleProblemController) Init(){

	//initalize
	s.clientChan = make(chan *problemControlRequest)
	s.problemChan = make(chan problems.ProblemUpdate)
	s.closeChan = make(chan bool)
	s.problemState = newSimpleProblemState()

	firstUpdate := s.problem.Init(s.problemChan)

	s.problemState.Alg = firstUpdate.Alg
	s.problemState.LastUpdate = firstUpdate.Data

	go s.problem.Loop()//Asynchronous execution of the problem loop (if it exists)

	go func() {
		var update = problems.ProblemUpdate{}
		var req = &problemControlRequest{}

		for{
			select{

			case update = <- s.problemChan:
				s.problemState.Update(update)

			case req = <- s.clientChan:
				if req.Data != nil{
					s.problemState.NewResult(req.Id, req.Data,
					 s.problem, req.Response)
				} else{
					s.problemState.NewRequest(req.Id, req.Response)
				}

			//close the goroutine	
			case <- s.closeChan:
				return		
			}
		}
	}()

}


func (s *SimpleProblemController) NewRequest(id uint32) (string, data, error){

	req := &problemControlRequest{id, nil, make(chan problemControlResponse)}
	
	s.clientChan <- req

	res := <- req.Response

	return res.Alg, res.Data, res.Err

}

func (s *SimpleProblemController) NewResult(id uint32, data []string) error{
	if data == nil{
		return errors.New("Result with no data")
	}

	req := &problemControlRequest{id, data, make(chan problemControlResponse)}
	
	s.clientChan <- req

	res := <- req.Response

	return res.Err
}

func (s *SimpleProblemController) Close(){
	s.closeChan <- true
}