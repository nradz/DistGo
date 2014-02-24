package channels

var clientChan = clientControlChannel{make(chan *ClientControlRequest)}
var problemChan = problemControlChannel{make(chan *ProblemControlRequest)}

func ClientControlChannel() *clientControlChannel{
	return &clientChan
}

func ProblemControlChannel() *ProblemControlChannel{
	return &problemChan
}

