package clientController

import(
	"testing"
	"time"
)

func TestClose(t *testing.T){
	c := New()
	c.Init()

	c.Close()
}

func TestNewClient(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make([]string,10)
	_, err := c.NewClient(header)

	if err != nil{
		t.Error(err.Error())
	} 
}

func TestLogged(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	time.Sleep(100 * time.Millisecond)

	header := make([]string, 10)
	id, _ := c.NewClient(header)


	err := c.IsLogged(id, header)

	if err != nil{
		t.Error(err.Error())
	}

}

func TestLoggedNotUserAgent(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make([]string, 10)
	id, _ := c.NewClient(header)

	otherHeader := make([]string, 11)
	err := c.IsLogged(id, otherHeader)

	if err == nil{
		t.Fail()
	}
}

func TestNotLogged(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make([]string, 10)
	
	err := c.IsLogged(10, header)

	if err == nil{
		t.Fail()
	}

}

func TestDeletedClient(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make([]string,10)
	id, _ := c.NewClient(header)

	err := c.DeleteClient(id, header)

	if err != nil{
		t.Error(err.Error())
	}
}

func TestNotDeletedClient(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make([]string,10)

	err := c.DeleteClient(10, header)

	if err == nil{
		t.Fail()
	}
}

func TestNotInitialized(t *testing.T){
	c := New()

	header := make([]string, 10)

	id, err := c.NewClient(header)
	if err == nil{
		t.Fatal("NewClient")
	}
	
	err = c.IsLogged(id,header)
	if err == nil{
		t.Fatal("IsLogged")
	}

	err = c.DeleteClient(id,header)
	if err == nil{
		t.Fatal("DeleteClient")
	}

	err = c.Close()
	if err == nil{
		t.Fatal("Close")
	}
}

