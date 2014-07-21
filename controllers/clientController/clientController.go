//Package clientController implements a controller to
//manage the clients of the system.
package clientController

import(
	//"github.com/nradz/DistGo/conf"
	"errors"
)

const(
	notStartedError string = "clientController is not started!" //the error 
	//that will be returned if clientController is not started.
)

type ClientController struct{
	started bool //started is true when clientController is running.
	clientChan chan *clientControlRequest //it is used to send data.
	//from the functions to the main loop.
	closeChan chan bool //It is used to finish the main loop.
	clist *clientList //It is the struct where the clients are saved.
}


//Request codes:
//10: New client
//20: isLogged
//30: Delete Client
type clientControlRequest struct{	
	Id uint32 //The id of the client
	Key uint32 //The key of the client
	Code uint8 //Code determines the type of the request
	UserAgent []string	//The userAgent of the client.
	Response chan clientControlResponse //The channel where the controller
	//will response.
}

type clientControlResponse struct{
	Id uint32 //The id of the client. Currently, it is only important
	//when the request was by newClient function
	Key uint32 //The key of the client. Currently, it is only important
	//when the request was by newClient function
	Err error
}

//return a new clientController struct.
func New() *ClientController{
	return &ClientController{}
}

//Init Initializes in another goroutine the main loop that manages
//the clients.
func (c *ClientController) Init(){
	
	//initialize
	c.clientChan = make(chan *clientControlRequest)
	c.closeChan = make(chan bool)
	c.clist = ClientList()

	go func(){
	
		var req = &clientControlRequest{}
		var res = clientControlResponse{}

		for {
			select{
			
			case req = <- c.clientChan:
				
				switch req.Code{
					case 10:
						res.Id, res.Key, res.Err = c.clist.newClient(req.UserAgent)
					case 20:
						res.Err = c.clist.isLogged(req.Id, req.Key, req.UserAgent)
					case 30:
						res.Err = c.clist.deleteClient(req.Id, req.Key, req.UserAgent)
				}

				req.Response <- res

			//close the goroutine	
			case <- c.closeChan:
				return

			}

		}

	}()

	c.started = true

}

//NewClient saves a new client and returns his "id". userAgent is used 
//to reduce the possibility of a phishing attack.
func (c *ClientController) NewClient(header map[string][]string) (uint32, uint32, error){
	if !c.started{
		return 0, 0, errors.New(notStartedError)
	}

	req := &clientControlRequest{0, 0, 10, header["User-Agent"], make(chan clientControlResponse)}

	c.clientChan <- req

	res := <- req.Response

	return res.Id, res.Key, res.Err
}

//IsLogged checks if a client is registered. It is true if error is "nil".
func (c *ClientController) IsLogged(id uint32, key uint32, header map[string][]string) error{
	if !c.started{
		return errors.New(notStartedError)
	}

	req := &clientControlRequest{id, key, 20, header["User-Agent"], make(chan clientControlResponse)}

	c.clientChan <- req

	res := <- req.Response

	return res.Err

}

//DeleteClient removes a client from the system.
func (c *ClientController) DeleteClient(id uint32, key uint32, header map[string][]string) error{
	if !c.started{
		return errors.New(notStartedError)
	}

	req := &clientControlRequest{id, key, 30, header["User-Agent"], make(chan clientControlResponse)}

	c.clientChan <- req

	res := <- req.Response

	return res.Err

}

//Close finishes the main loop.
func (c *ClientController) Close() error{
	if !c.started{
		return errors.New(notStartedError)
	}

	c.started = false
	c.closeChan <- true

	return nil
}