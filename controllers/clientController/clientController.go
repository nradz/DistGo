package clientController

import(
	//"fmt"
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
					id, t = newClient(creq.Header)
				case 20:
					id, t = isLogged(creq.Id, creq.Header)

				default:
					id, t = unknownError(creq.Id, creq.Header)

			}

			cres := channels.ClientControlResponse{id,t}

			creq.Response <- cres

		}

	}

	


}


func newClient(header map[string][]string) (uint32, uint8){

	var idres uint32 = clist.newClient(header)

	var tRes uint8 = 10

	return idres, tRes
}

func isLogged(id uint32, header map[string][]string) (uint32, uint8){

	cSaved, err := clist[id]

	var tRes uint8 = 0

	eq := reflect.DeepEqual(header, cSaved.Header)


	switch{
		case !err:
			tRes = 1

		case eq:
			tRes = 2 

		default:
			tRes = 20		

	}

	return id, tRes

}

func unknownError(id uint32, header map[string][]string) (uint32, uint8){

	return id, 0

}