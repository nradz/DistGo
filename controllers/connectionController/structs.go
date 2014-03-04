package connectionController

type ServerResponse struct{
Id uint32
Code uint8
Alg string
Data []string
}

type ClientRequest struct{
Id uint32
Code uint8
Data []string
}