package clientController

import(
	"fmt"
	"reflect"
	"github.com/nradz/DistGo/channels"	
	"github.com/nradz/DistGo/configuration"
)

var (
	clientChan = channels.ClientControlChannel()
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


func newClient(userAgent []string) (uint32, uint8){

	var idres uint32 = clist.newClient(userAgent)

	var tRes uint8 = 10

	fmt.Println("New Client:", idres)

	fmt.Println(clist)

	return idres, tRes
}


func isLogged(id uint32, userAgent []string) (uint32, uint8){

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

func deleteClient(id uint32, userAgent []string) (uint32, uint8){

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


func unknownError(id uint32, userAgent []string, code uint8) (uint32, uint8){
	fmt.Println("Error-> Code: ",code)
	return id, 0

}