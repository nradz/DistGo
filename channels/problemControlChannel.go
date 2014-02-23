package channels

type problemControlChannel struct{
	mainChan chan *ProblemControlRequest
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
	Data []string
}