package channels

var (
	clientChan = clientControlChannel{make(chan *ClientControlRequest)}
	problemChan = problemControlChannel{make(chan *ProblemControlRequest),
		make(chan ProblemUpdate)}
	)

func ClientControlChannel() *clientControlChannel{
	return &clientChan
}

func ProblemControlChannel() *problemControlChannel{
	return &problemChan
}

