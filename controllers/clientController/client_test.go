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

	header := make(map[string][]string)
	_, _, err := c.NewClient(header)

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
	id, key, _ := c.NewClient(header)


	err := c.IsLogged(id, key, header)

	if err != nil{
		t.Error(err.Error())
	}

}

func TestLoggedNotHeader(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make(map [string][]string)
	id, key, _ := c.NewClient(header)

	otherHeader := make(map [string][]string)
	otherheader["ala"] = nil
	err := c.IsLogged(id, key, otherHeader)

	if err == nil{
		t.Fail()
	}
}

func TestLoggedNotKey(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make(map [string][]string)
	id, key, _ := c.NewClient(header)

	err := c.IsLogged(id, 76543, header)
}

func TestNotLogged(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make(map [string][]string)
	
	err := c.IsLogged(10, 876543, header)

	if err == nil{
		t.Fail()
	}

}

func TestDeletedClient(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make(map[string][]string)
	id, key, _ := c.NewClient(header)

	err := c.DeleteClient(id, key, header)

	if err != nil{
		t.Error(err.Error())
	}
}

func TestNotDeletedClient(t *testing.T){
	c := New()
	c.Init()
	defer c.Close()

	header := make(map[string][]string)

	err := c.DeleteClient(10, 76543, header)

	if err == nil{
		t.Fail()
	}
}

func TestNotInitialized(t *testing.T){
	c := New()

	header := make([]string, 10)

	id, key, err := c.NewClient(header)
	if err == nil{
		t.Fatal("NewClient")
	}
	
	err = c.IsLogged(id, key, header)
	if err == nil{
		t.Fatal("IsLogged")
	}

	err = c.DeleteClient(id, key, header)
	if err == nil{
		t.Fatal("DeleteClient")
	}

	err = c.Close()
	if err == nil{
		t.Fatal("Close")
	}
}

