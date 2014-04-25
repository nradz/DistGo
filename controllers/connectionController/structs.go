package connectionController

type data interface{}

type ServerResponse struct{
Id uint32
Code uint8
Alg string
Data data
}

type ClientRequest struct{
Id uint32
Code uint8
Data []string
}