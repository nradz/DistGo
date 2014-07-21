package clientController

import(
	"reflect"
	"math/rand"
	"time"
	"errors"
	"github.com/nradz/DistGo/conf"
)

//The clienList type is a map where the clients are indexed 
//by their "id".
type clientList struct{
	nextId uint32 //The next position to the last one used
	list []client //The list where the clients are saved
	numClients uint32 //The amount of clients that are registered.
}

//The client struct
type client struct{
	key uint32 //A key to verify the client.
	userAgent []string //The user agent of the client.
	occupied bool //True if the position is been used
}


//ClientList returns a struct where the clients will be managed.
func ClientList() *clientList{

	cl := clientList{}

	cl.nextId = 0
	cl.list = make([]client, conf.MaxClients()) //The maximum number of clients
	//is determined in the configuration file.
	cl.numClients = 0

	return &cl
}

//newClient saves a new client into the struct. userAgent is used 
//to reduce the possibility of a phishing attack.
func (cl *clientList) newClient(userAgent []string) (uint32, uint32, error){
	if cl.numClients >= conf.MaxClients(){
		return 0, 0, errors.New("It have reached the maximum number of clients")
	}

	//Search an available position in the clientLst
	var id uint32 = cl.nextId
	var found bool = false
	for !found{
		switch{
		
		case id >= conf.MaxClients(): //If the 'index' is bigger than 
		//the maximum number of clients, it will be reset to zero
			id = 0

		case cl.list[id].occupied: //If the position is not available
			id += 1

		default: //The position is available
			cl.nextId = id + 1 //It prepares nextId to the next time
			found = true
		}

	}

	//Use a random number as key
	rand.Seed(time.Now().UTC().UnixNano())
	var key uint32 = rand.Uint32()

	var cli client = client{key, userAgent, true}


	//Add the client to the list
	cl.list[id] = cli 
	cl.numClients += 1

	return id, key, nil
}

//IsLogged checks if a client is registered. It is true if error is "nil".
func (cl *clientList) isLogged(id uint32, key uint32, userAgent []string) error{
	
	if id >= conf.MaxClients(){
		return errors.New("Id is not valid")
	}

	client := cl.list[id]

	if !client.occupied{ //The position is free
		return errors.New("Client does not exist")

	} else if client.key != key{
		return errors.New("Incorrect key")

	} else if eq := reflect.DeepEqual(userAgent, client.userAgent); !eq {
		return errors.New("UserAgent not equal")

	} else{
		//Client is logged
		return nil
	}

}

//deleteClient removes a client from the struct
func (cl *clientList) deleteClient(id uint32, key uint32, userAgent []string) error{

	if err := cl.isLogged(id, key, userAgent); err != nil{
		return err

	}else{
		cl.list[id].occupied = false
		cl.numClients -= 1
		return nil

	}
}