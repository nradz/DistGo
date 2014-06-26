package problemController

import(
	"fmt"
	"github.com/nradz/DistGo/configuration"
	"github.com/nradz/DistGo/problems"
	)

type data interface{}

type ProblemControlRequest struct{
	Id uint32
	Data []string
	Response chan ProblemControlResponse
}

type ProblemControlResponse struct{
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

	problemState.LastUpdate = nil
	problemState.Clients = make(map[uint32]*clientState)

	firstUpdate := problem.Init(problemChan)

	problemState.Alg = firstUpdate.Alg
	problemState.LastUpdate = firstUpdate.Data

	go problem.Loop()//Asynchronous execution of the problem loop (if it exists)


	var update = problems.ProblemUpdate{}
	var req = &ProblemControlRequest{}
	var res = ProblemControlResponse{}

	for{
		select{
		case update = <-problemChan:
			problemState.Update(update)

		case req = <-clientChan:
			if req.Data != nil{
				res.Err = problemState.NewResult(req.Id, req.Data)
			} else{
				res.Alg, res.Data, res.Err = problemState.NewRequest(req.Id)
			}

			req.Response <- res

		}
	}

}


func NewRequest(id uint32){

	req := &ProblemControlRequest{id, nil, make(chan ProblemControlResponse)}
	
	clientChan <- req

	res := <- req.Response

	return res.Alg, res.Data, res.Err

}

func NewResult(id uint32, data []string){
	req := &ProblemControlRequest{id, data, make(chan ProblemControlResponse)}
	
	clientChan <- req

	res := <- req.Response

	return res.Err
}

/////////fromClient and his functions///////////////

func fromClient(req *channels.ProblemControlRequest, problem problems.Problem){
	cState, ok := problemState.Clients[req.Id]
	switch{
		// New client
		case !ok && (req.Code == 20):
			newClient(req)
		// Standard Request
		case ok && (req.Code == 20):
			newRequest(req, cState)
		// Request with new result
		case ok && (req.Code == 30):
			//newResult(req, cState)
			newResult(req, cState)
			//Pass the data to the problem (asynchronous)
			go problem.NewResult(req.Data)

		default:
			fmt.Println("Error in simpleProblemController.fromClient-> Code:%d OK:%t", req.Code, ok)
	}

}




func newClient(req *channels.ProblemControlRequest){
	cState := clientState{false, true, nil}

	problemState.Clients[req.Id] = &cState

	//Make the response with the algorithm and the last update
	res := channels.ProblemControlResponse{130, problemState.Alg, problemState.LastUpdate}

	req.Response <- res
}

func newRequest(req *channels.ProblemControlRequest, cState *clientState){

	switch{

	//Client in standby
	case cState.Updated && !cState.Ready:
		cState.Ready = true
		cState.ResChan = req.Response

	case !cState.Updated && !cState.Ready:
		res := channels.ProblemControlResponse{120, "", problemState.LastUpdate}
		req.Response <-res

	default:
		fmt.Println("Error in simpleProblemController.newRequest-> Updated: %t Ready: %t",
			cState.Updated, cState.Ready)
	}

}


func newResult(req *channels.ProblemControlRequest, cState *clientState){

	res := channels.ProblemControlResponse{140, "", nil}
	req.Response <- res
}

/////////fromProblem and his functions///////////////

func fromProblem(update channels.ProblemUpdate){

	problemState.LastUpdate = update.Data

	for _, cState := range problemState.Clients{

		switch{
		
		case cState.Ready && cState.Updated:
			cState.Ready = false
			cState.ResChan <- channels.ProblemControlResponse{120, "", problemState.LastUpdate}

		case !cState.Ready && cState.Updated:
			cState.Updated = false

		default:
		fmt.Println("Error in simpleProblemController.fromProblem-> Updated:",
			cState.Updated, " Ready: ", cState.Ready)

		}
	}

}