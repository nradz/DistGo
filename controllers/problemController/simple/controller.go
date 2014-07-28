package simple

import(
	//"fmt"
	"errors"
	//"github.com/nradz/DistGo/conf"
	"github.com/nradz/DistGo/problems"
	)

type Simple struct{
	clientChan chan *request //It is used 
	//to send data from the functions to the main loop
	problemChan chan problems.ProblemUpdate //It is used
	//to send data from the problem to the main loop
	standbyChan chan *request //If a client is updated, he will 
	//be set in standby.
	closeChan chan bool //It is used to finish the main loop
	problem problems.Problem //The problem that is been executed
	alg string //The actual algorithm that is executing in the clients
	updateData problems.Data //The last update available
	updateNumber uint32 //The number of the update
}

//It is the struct used to send data from the functions 
//to the main loop
type request struct{
	Id uint32 //Id of the client
	Data []string //Data of the client (If the request is 
		//a new result)
	LastNumber uint32 //The number of the last update that
	//the client has received
	Response chan response //The chan where the main loop will 
	//response to the functions
}

//The response struct
type response struct{
	Alg string //The client algorithm
	Data problems.Data //The update data
	Number uint32 //Number of the update
	Err error
}

//Return a new simple problem controller
func New(prob problems.Problem) *Simple{
	return &Simple{
		problem: prob,
	}

}

//it initializes the controller and starts the main loop
func (s *Simple) Init(){

	//initalize
	s.clientChan = make(chan *request)
	s.problemChan = make(chan problems.ProblemUpdate)
	s.closeChan = make(chan bool)
	s.standbyChan = make(chan *request)

	firstUpdate := s.problem.Start(s.problemChan)

	s.alg = firstUpdate.Alg
	s.updateData = firstUpdate.Data
	s.updateNumber = firstUpdate.Number

	go s.problem.Loop() //Asynchronous execution of the problem 
	//loop (if it exists)

	go func() {
		var update = problems.ProblemUpdate{}
		var req = &request{}

		for{
			select{

			case update = <- s.problemChan:
				s.fromProblem(update)

			case req = <- s.clientChan:
				s.fromClient(req)

			//close the goroutine	
			case <- s.closeChan:
				return		
			}
		}
	}()

}

//public functions

//NewRequest sends a new request to the controller. It returns 
//the client algorithm, update data, update number and the error.
//Algorithm always is 'Zero', except when 'lastUpdate' 
//is zero (the first request).
func (s *Simple) NewRequest(id uint32, lastUpdate uint32) (string, problems.Data, uint32, error){

	req := &request{id, nil, lastUpdate, make(chan response)}
	
	s.clientChan <- req

	res := <- req.Response

	if res.Err != nil{
		return "", nil, 0, res.Err
	}

	if res.Number == lastUpdate{ //The client is updated. So, it will be set in 'standby'.
		s.standbyChan <- req
		res = <- req.Response
	}

	return res.Alg, res.Data, res.Number, res.Err

}

//NewResult sends a new result to the controller.
func (s *Simple) NewResult(id uint32, data []string, lastUpdate uint32) error{
	if data == nil{
		return errors.New("Result with no data")
	}

	go s.problem.NewResult(data, lastUpdate)//It executes asynchronously the 
	//NewResult function of the problem

	return nil
}

//Close stops the main loop of the controller.
func (s *Simple) Close(){
	s.closeChan <- true
}


//private functions

//fromProblem manages 
func (s *Simple) fromProblem(update problems.ProblemUpdate){
	s.updateData = update.Data
	s.updateNumber = update.Number

	//Update the clients that were in standby
	go func(update problems.Data, number uint32){
		for{
			select{
			case req := <- s.standbyChan:
				req.Response <- response{"", update, number, nil}
			default:
				return
			}
		}
	}(update.Data, update.Number)
}

func (s *Simple) fromClient(req *request){
		
	switch{
	case req.LastNumber == 0: //The client needs the algorithm.
		req.Response <- response{s.alg, s.updateData, s.updateNumber, nil}

	case req.LastNumber < s.updateNumber: //The client has an old update.
		req.Response <- response{"", s.updateData, s.updateNumber, nil}
		
	case req.LastNumber == s.updateNumber: //The client is updated.
		req.Response <- response{"", nil, s.updateNumber, nil}
		
	default:
		req.Response <- response{"", nil, s.updateNumber, 
			errors.New("The update number of the client is bigger")}
	}

}