package channels

type clientControlChannel struct{
	mainChan chan *ClientControlRequest
}

type ClientControlRequest struct{	
	Id uint32
	Code uint8
	Header map[string][]string	
	Response chan ClientControlResponse
}

type ClientControlResponse struct{
	Id uint32
	Code uint8
}





//////////////////////////////////////////////////

func (clientChan *clientControlChannel) Send(request *ClientControlRequest) ClientControlResponse{

	clientChan.mainChan <- request

	return <-request.Response

}


func (clientChan *clientControlChannel) Receive() (chan *ClientControlRequest){

	return clientChan.mainChan
}