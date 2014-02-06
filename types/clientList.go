package distgo_types

import(
	"math/rand"
	"time"
)

type ClientList struct{
	List map[uint32]Client
}

type Client struct{
	Ip string
}


//Generate a new ClientList with default values
func NewClientList() ClientList{

	cl := ClientList{}

	cl.List = make(map[uint32]Client)

	return cl
}

func (l *ClientList) NewClient(ip string){

	//Use a random number as a id and check if it is available.
	rand.Seed(time.Now().UTC().UnixNano())
	var id uint32 = rand.Uint32()
	
	for used := true; used == true;{
		_, used = l.List[id]
		if used{
			rand.Seed(time.Now().UTC().UnixNano())
			id = rand.Uint32()
		}
	}

	var cli Client = Client{ip}



	//Add the client to the list
	l.List[id] = cli 

}