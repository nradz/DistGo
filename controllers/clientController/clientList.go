package clientController

import(
	"reflect"
	"math/rand"
	"time"
	"errors"
)

type clientList map[uint32]client

type client struct{
	userAgent []string
}


//Generate a new ClientList with default values
func ClientList() clientList{

	cl := clientList{}

	cl = make(map[uint32]client)

	return cl
}

func (l clientList) newClient(userAgent []string) (uint32, error){

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

	var cli client = client{userAgent}


	//Add the client to the list
	l[id] = cli 

	return id, nil
}

func (l clientList) isLogged(id uint32, userAgent []string) error{
	
	saved, ok := l[id]

	eq := reflect.DeepEqual(userAgent, saved.userAgent)

	switch{
	//Id not found
	case !ok:
		return errors.New("Id not found")
	//UserAgent not equal
	case !eq:
		return errors.New("UserAgent not equal")
	//Logged Client
	default:
		return nil
	}

}

func (l clientList) deleteClient(id uint32, userAgent []string) error{

	if err := l.isLogged(id, userAgent); err != nil{
		return err
	}else{
		delete(l, id)
		return nil
	}
}