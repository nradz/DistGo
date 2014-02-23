package channels

var clientChan = clientControlChannel{make(chan *ClientControlRequest)}

func ClientControlChannel() *clientControlChannel{
	return &clientChan
}

