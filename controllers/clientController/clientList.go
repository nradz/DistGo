package clientController

import(
	"reflect"
	"math/rand"
	"time"
	"errors"
)

//The clienList type is a map where the clients are indexed 
//by their "id".
type clientList map[uint32]client

//The client struct
type client struct{
	userAgent []string //The user agent of the client.
}


//ClientList returns a struct where the clients will be managed.
func ClientList() clientList{

	cl := clientList{}

	cl = make(map[uint32]client)

	return cl
}

//newClient saves a new client into the struct. userAgent is used 
//to reduce the possibility of a phishing attack.
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

//IsLogged checks if a client is registered. It is true if error is "nil".
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

//deleteClient removes a client from the struct
func (l clientList) deleteClient(id uint32, userAgent []string) error{

	if err := l.isLogged(id, userAgent); err != nil{
		return err
	}else{
		delete(l, id)
		return nil
	}
}