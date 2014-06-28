package connectionController

import(
	"testing"
)


func TestNotPostMethod(t *testing.T){
	sr, cr, head := setup()

	notPostMethod(sr, cr, head)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}

}

func TestNotValidJSON(t *testing.T){
	sr, cr, head := setup()

	notValidJSON(sr, cr, head)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}

}

func TestNotValidCode(t *testing.T){
	sr, cr, head := setup()

	notValidCode(sr, cr, head)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}
}

func TestNewClient(t *testing.T){
	sr, cr, head := setup()

	newClient(sr, head)

	if sr.Id == 0{
		t.Error("No id:", sr.Id)
	}

	if sr.Code != 110{
		t.Error("Incorrect Code: ", sr.Code)
	}

	if sr.Alg != ""{
		t.Error("Alg: ", sr.Alg)
	}

	if data != nil{
		t.Error("Data: ", sr.data)
	}

}

func TestFirstRequest(t *testing.T){
	sr, cr, head := setup()

	newClient(sr, head)

	cr.Id = sr.Id

	newRequest(sr, cr.Id, cr.Head)

	if sr.Id != cr.Id{
		t.Error("No Id:" sr.Id)
	}

	if sr.
	
}


func setup() (*ServerResponse, ClientRequest, []string) {
	sr := &ServerResponse{}
	cr := ClientRequest{0, 10, nil}
	head := make([]string)
}