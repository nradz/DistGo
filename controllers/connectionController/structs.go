package connectionController

//Generic data
type data interface{}

//ServerResponse is the struct that is used by
//the server to send data to the clients.
/*	Code:
	0   -> Error
	110 -> New Id and key (response to "new client")
	120 -> New data (update)
	130 -> New algorithm (currently, the response of the first request)
	140 -> New result received
	150 -> Client deleted (connection end)*/
type ServerResponse struct{
	Id uint32
	Key uint32
	Code uint8
	Alg string
	Status uint32
	Data data
}

//ClientRequest is the struct that is used by
//the clients to send data to the server.
/*	Code:
	10 -> new client
	20 -> new request
	30 -> new result
	100-> delete client*/
type ClientRequest struct{
	Id uint32
	Key uint32
	Code uint8
	LastUpdate uint32
	Data []string
}