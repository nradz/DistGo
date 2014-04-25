package channels

type data interface{}

type problemControlChannel struct{
	reqChan chan *ProblemControlRequest
	updateChan chan ProblemUpdate

}

type ProblemControlRequest struct{	
	Id uint32
	Code uint8
	Data []string
	Response chan ProblemControlResponse
}

type ProblemControlResponse struct{
	Code uint8
	Alg string
	Data data
}

type ProblemUpdate struct{
	Alg string
	Data data
}





///////////////////////////////////
func (problemChan *problemControlChannel) SendRequest(request *ProblemControlRequest) ProblemControlResponse{

	problemChan.reqChan <- request

	return <-request.Response

}


func (problemChan *problemControlChannel) ReceiveRequest() (chan *ProblemControlRequest){

	return problemChan.reqChan
}

func (problemChan *problemControlChannel) SendUpdate(update ProblemUpdate){

	problemChan.updateChan <- update
}

func (problemChan *problemControlChannel) ReceiveUpdate() (chan ProblemUpdate){

	return problemChan.updateChan
}