package clientController

import(
	"fmt"
	"reflect"
	"error"
	"github.com/nradz/DistGo/channels"	
	"github.com/nradz/DistGo/configuration"
)

type ClientControlRequest struct{	
	Id uint32
	Code uint8
	UserAgent []string	
	Response chan ClientControlResponse
}

type ClientControlResponse struct{
	Id uint32
	Code uint8
}

var (
	clientChan = make(chan *ClientControlRequest)
	conf = configuration.Configuration()
	)

var clist ClientList = ClientList{}

func ClientController(){

	clist = newClientList()

	var id uint32 = 0
	var t uint8 = 0

	for {
		select{
		
		case creq := <- clientChan.Receive():
			
			switch creq.Code{
				case 10:
					id, t = newClient(creq.UserAgent)
				case 20:
					id, t = isLogged(creq.Id, creq.UserAgent)
				case 30:
					id, t = deleteClient(creq.Id, creq.UserAgent)

				default:
					id, t = unknownError(creq.Id, creq.UserAgent, creq.Code)

			}

			cres := channels.ClientControlResponse{id,t}

			creq.Response <- cres

		}

	}

}


func NewClient(userAgent []string) (uint32, error){

	var idres uint32 = clist.newClient(userAgent)

	fmt.Println("New Client:", idres)

	return idres, nil
}


func IsLogged(id uint32, userAgent []string) (uint32, error){

	cSaved, ok := clist[id]

	var tRes uint8 = 0

	eq := reflect.DeepEqual(userAgent, cSaved.UserAgent)

	switch{
	case !ok:
		tRes = 1

	case !eq:
		tRes = 2 

	default:
		tRes = 20		

	}

	return id, tRes

}

func DeleteClient(id uint32, userAgent []string) (uint32, error){

	cSaved, ok := clist[id]
	
	var tRes uint8 = 0

	eq := reflect.DeepEqual(userAgent, cSaved.UserAgent)

	switch{
	case !ok:
		tRes = 1

	case !eq:
		tRes = 2

	default:
		delete(clist, id)
		tRes = 30
	}

	fmt.Println(clist)

	return id, tRes
}