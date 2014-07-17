package clientController

import(
	"reflect"
	"math/rand"
	"time"
	"errors"
	"github.com/nradz/conf"
)

//The clienList type is a map where the clients are indexed 
//by their "id".
type clientList struct{
	nextId uint16 //The next position to the last one used
	list []client //The list where the clients are saved
}

//The client struct
type client struct{
	key uint32 //A key to verify the client.
	userAgent []string //The user agent of the client.
}


//ClientList returns a struct where the clients will be managed.
func ClientList() clientList{

	cl := clientList{}

	cl.nextId = 0
	cl.list = make([]client, conf.NClients()) //The maximum number of clients
	//is determined in the configuration file.

	return cl
}

//newClient saves a new client into the struct. userAgent is used 
//to reduce the possibility of a phishing attack.
func (cl clientList) newClient(userAgent []string) (uint32, error){

	//Search an available position in the clientLst
	var id uint16 = cl.nextId
	for{
		switch{
		
		case id > conf.NClients(): //If the 'index' is bigger than 
		//the maximum number of clients, it will be reset to zero
			id = 0

		case cl.list[id] != nil: //If the position is not available
			id += 1

		default: //The position is available
			cl.nextId = id + 1 //It prepares nextId to the next time
			break
		}

	}

	//Use a random number as key
	rand.Seed(time.Now().UTC().UnixNano())
	var key uint32 = rand.Uint32()

	var cli client = client{key, userAgent}


	//Add the client to the list
	cl.list[id] = cli 

	return id, key, nil
}

//IsLogged checks if a client is registered. It is true if error is "nil".
func (cl clientList) isLogged(id uint16, key uint32, userAgent []string) error{
	
	if id < conf.NClients(){
		return errors.New("Id is not valid")
	}

	client := cl.list[id]

	if client == nil{
		return errors.New("Client does nott exist")

	} else if client.key != key{
		return errors.New("Incorrect key")

	} else if eq := reflect.DeepEqual(userAgent, saved.userAgent); !eq {
		return errors.New("UserAgent not equal")

	} else{
		//Client is logged
		return nil
	}

}

//deleteClient removes a client from the struct
func (cl clientList) deleteClient(id uint32, key, userAgent []string) error{

	if err := l.isLogged(id, key, userAgent); err != nil{
		return err
	}else{
		cl.list[id] = nil
		return nil
	}
}