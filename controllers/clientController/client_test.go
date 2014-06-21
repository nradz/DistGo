package clientController

import(
	"testing"
)


func TestNewClient(t *testing.T){
	go ClientController()
	header := make([]string,10)
	_, err := NewClient(header)

	if err != nil{
		t.Error(err.Error())
	} 
}

func TestLogged(t *testing.T){
	go ClientController()
	header := make([]string, 10)
	id, _ := NewClient(header)


	err := IsLogged(id, header)

	if err != nil{
		t.Error(err.Error())
	}

}

func TestLoggedNotUserAgent(t *testing.T){
	go ClientController()
	header := make([]string, 10)
	id, _ := NewClient(header)

	otherHeader := make([]string, 11)
	err := IsLogged(id, otherHeader)

	if err == nil{
		t.Fail()
	}
}

func TestNotLogged(t *testing.T){
	go ClientController()
	header := make([]string, 10)
	
	err := IsLogged(10, header)

	if err == nil{
		t.Fail()
	}

}

func TestDeletedClient(t *testing.T){
	go ClientController()
	header := make([]string,10)
	id, _ := NewClient(header)

	err := DeleteClient(id, header)

	if err != nil{
		t.Error(err.Error())
	}
}

func TestNotDeletedClient(t *testing.T){
	go ClientController()
	header := make([]string,10)

	err := DeleteClient(10, header)

	if err == nil{
		t.Fail()
	}
}


