package problemController

import(
//"fmt"
"errors"
"github.com/nradz/DistGo/problems"
"github.com/nradz/DistGo/conf"
)

type simpleProblemState struct{
	Alg string //The actual algorithm that is executing in the clients
	LastUpdate data //The last update available
	Clients []*clientState //Map of the clients with their state
}

type clientState struct{
	Ready bool //A bool variable that indicate if the client is ready for a update
	Updated bool //It indicate if the client have received the last update
	ResChan chan problemControlResponse /*The channel where the problemController
	receive the request*/
}

func newSimpleProblemState() *simpleProblemState {

	sp := &simpleProblemState{
		Clients: make([]*clientState, conf.NClients()),
	}

	return sp
}

func (sp *simpleProblemState) NewRequest(id uint32, res chan problemControlResponse){
	client, ok := sp.Clients[id]
	var alg string = ""

	if !ok{
		client = &clientState{false, false, nil}
		sp.Clients[id] = client
		//the first time that the client get the algorithm
		alg = sp.Alg
	}

	switch{
	//Client will be set in standby
	case client.Updated && !client.Ready:
		client.Ready = true
		client.ResChan = res
		
	//Client will be updated	
	case !client.Updated && !client.Ready:
		client.Updated = true
		res <- problemControlResponse{alg, sp.LastUpdate, nil}
		
	default:
		err := errors.New("Unknown Error")
		res <- problemControlResponse{"", nil, err}
	}
}


func (sp *simpleProblemState) NewResult(id uint32, result []string, 
  prob problems.Problem, res chan problemControlResponse){

  	var err error = nil

  	//If the client didn't a first request, he dont have the algorithm.
  	//So, he can't send a valid result.
  	if _, ok := sp.Clients[id]; !ok{
  		err = errors.New("The client do not have the algorithm! He can't send a valid result.")
  	} else{
		//Pass the data to the problem (asynchronous)
		go prob.NewResult(result)
	}

	res <- problemControlResponse{"", nil, err}

}

func (sp *simpleProblemState) Update(update problems.ProblemUpdate){
	sp.LastUpdate = update.Data

	for _, client := range sp.Clients{
		switch{
		//Client is in standby
		case client.Ready && client.Updated:
			client.Ready = false
			client.ResChan <-problemControlResponse{"", sp.LastUpdate, nil}
			
		
		case !client.Ready && client.Updated:
			client.Updated = false
			

		}
	}
}