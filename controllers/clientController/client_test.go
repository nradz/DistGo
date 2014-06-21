package clientController

import(
	"testing"
)


func TestNewClient(t *testing.T){
	go ClientController()
	header := make([]string,10)
	id, err := NewClient(header)

	if err != nil{
		t.Fail()
	} 
}

func TestLogged(t *testing.T){
	go ClientController()
	header := make([]string, 10)
	id, err := NewClient(header)


	logged, err := IsLogged(id, header)

	if logged != true{
		t.Fail()
	}

}

func TestNotLogged(t *testing.T){
	go ClientController()
	header := make([]string, 10)
	
	logged, err := IsLogged(10, header)

	if logged != false{
		t.Fail()
	}

}

func TestDeletedClient(t *testing.T){
	go ClientController()
	header := make([]string,10)
	id, err := NewClient(header)

	deleted, err := DeleteClient(id, header)

	if deleted != true{
		t.Fail()
	}
}

func TestNotDeletedClient(t *testing.T){
	go ClientController()
	header := make([]string,10)

	deleted, err := DeleteClient(10, header)

	if deleted != false{
		t.Fail()
	}
}


