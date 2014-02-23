package clientController

import(
	"math/rand"
	"time"
)

type ClientList map[uint32]Client

type Client struct{
	Header map[string][]string
}


//Generate a new ClientList with default values
func newClientList() ClientList{

	cl := ClientList{}

	cl = make(map[uint32]Client)

	return cl
}

func (l ClientList) newClient(header map[string][]string) uint32{

	//Use a random number as a id and check if it is available.
	rand.Seed(time.Now().UTC().UnixNano())
	var id uint32 = rand.Uint32()
	
	for used := true; used == true;{
		_, used = l[id]
		if used || (id == 0){
			rand.Seed(time.Now().UTC().UnixNano())
			id = rand.Uint32()
		}
	}

	var cli Client = Client{header}


	//Add the client to the list
	l[id] = cli 

	return id
}