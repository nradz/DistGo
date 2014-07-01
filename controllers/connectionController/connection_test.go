package connectionController

import(
	"http"
	"json"
	"testing"
	"clientController"
	"problemController"
	"problems"
)

var url string = "localhost:8080"

func TestNotPostMethod(t *testing.T){
	cc := setup()
	defer cc.Close()

	resp, err := http.Get("localhost:8080")
	if err != nil{
		t.Error("Error de conexion")
	}
	defer resp.Body.Close()

	sr, err := decode(res.Body)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}

}

func TestNotValidJSON(t *testing.T){
	cc := setup()
	defer cc.Close()

	resp, err 

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}

}

func TestNotValidCode(t *testing.T){
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()

	cc.NotValidCode(sr, cr, head)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}
}

func TestNewClient(t *testing.T){
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()

	cc.NewClient(sr, head)

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
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()

	cc.NewClient(sr, head)

	cr.Id = sr.Id

	cc.NewRequest(sr, cr, Head)

	if sr.Id != cr.Id{
		t.Error("No Id:" sr.Id)
	}

	if sr.Code !=130{
		t.Error("No Code:", sr.Code)
	}

	if sr.Alg == ""{
		t.Error("No alg:", sr.Alg)
	}

	if sr.Data == nil{
		t.Error("No data:", sr.Data)
	}

}


func TestNotLoggedRequest(t *testing.T){
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()

	cc.NewRequest(sr, cr, Head)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestIncorrectHeader(t *testing.T){
	cc := ConnectionController()
	cc.Init()
	defer cc.Close()

	sr, cr, head := setup()

	cc.NewClient(sr, head)

	cr.Id = sr.Id

	false_head := make(string[], 9)

	cc.NewRequest(sr, cr, false_head)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}

}

func TestNewResult(t *testing.T){
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()

	cc.NewClient(sr,head)

	cr.Id = sr.Id

	cc.NewRequest(sr, cr, head)

	cr.Data := make([]string,1)
	cr.Data[0] = "6"

	cc.NewResult(sr, cr, head)

	if sr.Code != 140{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNewResultNoAlg(t *testing.T){
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()

	cc.NewClient(sr,head)

	cr.Id = sr.Id

	cr.Data := make([]string,1)
	cr.Data[0] = "6"

	cc.NewResult(sr, cr, head)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNotLoggedResult(t *testing.T){
	cc := setup()
	defer cc.Close()

	sr, cr, head := datos()
}


func setup() *connectionController{
	
	cli := clientController.ClientController()
	prob := problems.GetProblem("pruebaProblem")
	prcon := problemController.ProblemController(prob)
	
	con := ConnectionController(cli, procon)

	con.Init()

	return con
}

func datos() (*ServerResponse, ClientRequest, []string){
	sr := &ServerResponse{}
	cr := ClientRequest{0, 10, nil}
	head := make([]string)
	return sr, cr, head
}


func decode(resp io.Reader) (*ServerResponse, error){

	dec := json.NewDecoder(resp)
	sr := &ServerResponse{}
	err:= dec.Decode(sr)

	return sr, err
}