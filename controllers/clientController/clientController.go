package clientController

import(
	"github.com/nradz/DistGo/configuration"
)


//Request codes:
//10: New client
//20: isLogged
//30: Delente Client
type clientControlRequest struct{	
	Id uint32
	Code uint8
	UserAgent []string	
	Response chan clientControlResponse
}

type clientControlResponse struct{
	Id uint32
	Err error
}

var (
	clientChan = make(chan *clientControlRequest)
	conf = configuration.Configuration()
	clist clientList =clientList{}
	)

func ClientController(){

	clist = newClientList()

	var id uint32 = 0
	var err error = nil

	for {
		select{
		
		case req := <- clientChan:
			
			switch req.Code{
				case 10:
					id, err = clist.newClient(req.UserAgent)
				case 20:
					err = clist.isLogged(req.Id, req.UserAgent)
				case 30:
					err = clist.deleteClient(req.Id, req.UserAgent)
			}

			res := clientControlResponse{id, err}

			req.Response <- res

		}

	}

}


func NewClient(userAgent []string) (uint32, error){

	req := &clientControlRequest{0, 10, userAgent, make(chan clientControlResponse)}

	clientChan <- req

	res := <- req.Response

	return res.Id, res.Err
}


func IsLogged(id uint32, userAgent []string) error{

	req := &clientControlRequest{id, 20, userAgent, make(chan clientControlResponse)}

	clientChan <- req

	res := <- req.Response

	return res.Err

}

func DeleteClient(id uint32, userAgent []string) error{

	req := &clientControlRequest{id, 30, userAgent, make(chan clientControlResponse)}

	clientChan <- req

	res := <- req.Response

	return res.Err

}