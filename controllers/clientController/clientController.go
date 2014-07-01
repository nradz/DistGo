package clientController

import(
	"github.com/nradz/DistGo/configuration"
)



type clientController struct{
	clientChan chan *clientControlRequest
	closeChan chan bool
	clist clientList
}


//Request codes:
//10: New client
//20: isLogged
//30: Delete Client
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

var conf = configuration.Configuration()

func ClientController() *clientController{
	return &clientController{}
}

func (c *clientController) Init(){
	
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
						res.Id, res.Err = c.clist.newClient(req.UserAgent)
					case 20:
						res.Err = c.clist.isLogged(req.Id, req.UserAgent)
					case 30:
						res.Err = c.clist.deleteClient(req.Id, req.UserAgent)
				}

				req.Response <- res

			//close the goroutine	
			case <- c.closeChan:
				return

			}

		}

	}()

}


func (c *clientController) NewClient(userAgent []string) (uint32, error){

	req := &clientControlRequest{0, 10, userAgent, make(chan clientControlResponse)}

	c.clientChan <- req

	res := <- req.Response

	return res.Id, res.Err
}


func (c *clientController) IsLogged(id uint32, userAgent []string) error{

	req := &clientControlRequest{id, 20, userAgent, make(chan clientControlResponse)}

	c.clientChan <- req

	res := <- req.Response

	return res.Err

}

func (c *clientController) DeleteClient(id uint32, userAgent []string) error{

	req := &clientControlRequest{id, 30, userAgent, make(chan clientControlResponse)}

	c.clientChan <- req

	res := <- req.Response

	return res.Err

}

func (c *clientController) Close(){
	c.closeChan <- true
}