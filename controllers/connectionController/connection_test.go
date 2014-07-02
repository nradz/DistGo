package connectionController

import(
	"http"
	"json"
	"testing"
	"clientController"
	"problemController"
	"problems"
)

var url string = "http://localhost:8080"
var json_type string = "application/json"

const(
	get_error = "The request method is not POST"
	json_error = "The JSON message is not valid"
	code_error = "Not valid Code"
	)

const(
	alg = `function mainFunc(romero,data){
				window.alert("miau");
					romero.finish();
				}
	`	
	)

func TestNotPostMethod(t *testing.T){
	cc := setup()
	defer cc.Close()

	resp, err := http.Get(url)
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()

	sr, err := decode(res.Body)
	if err != nil{
		t.Fatal(err)
	}

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
	if sr.Data.(string) != get_error{
		t.Error("Incorrect message:", sr.Data)
	}

}

func TestNotValidJSON(t *testing.T){
	cc := setup()
	defer cc.Close()

	req := bytes.NewBufferString("Hola, niño")

	resp, err := http.Post(url, "text", req)
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil{
		t.Fatal(err)
	}

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}
	if sr.Data.(string) != json_error{
		t.Error("Incorrect message:", sr.Data)
	}

}

func TestNotValidCode(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{324,nil,nil}
	
	sr := normalPost(cr, t)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}
	if sr.Data.(string) != code_error{
		t.Error("Incorrect message:", sr.Data)
	}
}

func TestNewClient(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 10, nil}

	sr := normalPost(cr, t)

	if sr.Id == 0{
		t.Error("No id:", sr.Id)
	}

	if sr.Code != 110{
		t.Error("Incorrect Code: ", sr.Code)
	}

}

func TestFirstRequest(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 10, nil}

	//login
	sr := normalPost(cr, t)

	cr.Id = sr.Id
	cr.Code = 20

	sr = normalPost(cr, t)

	if sr.Id != cr.Id{
		t.Error("No Id:" sr.Id)
	}

	if sr.Code !=130{
		t.Error("No Code:", sr.Code)
	}

	if sr.Alg != alg{
		t.Error("No alg:", sr.Alg)
	}

	if sr.Data.(int) == nil{
		t.Error("No data:", sr.Data)
	}

}


func TestNotLoggedRequest(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 20, nil}

	sr := normalPost(cr, t)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestIncorrectHeader(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 10, nil}

	//login
	sr := normalPost(cr, t)

	req, err := encode(cr)
	if err != nil{
		t.Fatal(err)
	}

	cr.Id = sr.Id
	cr.Code = 20

	json, err := encode(cr)
	if err != nil{
		t.Fatal(err)
	}

	client := &http.Client{}

	newr, err := http.NewRequest("POST", url, json)
	if err != nil{
		t.Fatal(err)
	}

	client.Do(newr)
	if err != nil{
		t.Fatal(err)
	}

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}

}

func TestNewResult(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 10, nil}

	sr := normalPost(cr, t)

	cr.Id = sr.Id
	cr.Code = 20
	
	sr = normalPost(cr, t)

	cr.Code = 30

	sr = normalPost(cr, t)

	if sr.Code != 140{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNewResultNoAlg(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 10, nil}

	sr := normalPost(cr, t)

	cr.Id = sr.Id
	cr.Code = 30

	sr = normalPost(cr, t)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNotLoggedResult(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := clientRequest{0, 30, nil}

	sr := normalPost(cr, t)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestUpdate(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr1 := client{0, 10, nil}
	cr2 := client{0, 10, nil}

	//login
	sr1 := normalPost(cr1, t)
	sr2 := normalPost(cr2, t)
	
	//first request
	cr1.Id = sr1.Id
	cr2.Id = sr2.Id
	cr1.Code = 20
	cr2.Code = 20	
	
	sr1 = normalPost(cr1, t)
	sr2 = normalPost(cr2, t)

	//result
	cr1.Code = 30
	cr1.Data = make([]string, 1)
	cr1.Data[0] = "6"

	sr1 = normalPost(cr1, t)

	//update
	cr2.Code = 20
	sr2 = normalPost(cr2, t)

	if sr2.Data.(int) != 6{
		t.Error("Incorrect data:", sr2.Data)
	}

}

func TestDelete(t *testing.T){
	cc := setup()
	defer cc.Close()

	cr := client{0, 10, nil}

	//login
	sr := normalPost(cr1, t)

	cr.Id = sr.Id
	cr.Code = 100

	sr = normalPost(cr, t)

	if sr.Code != 150{
		t.Fatal("Incorrect Code:", sr.Code)
	}

	//Verify

	cr.Code = 20
	sr = normalpost(cr, t)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestDeteleNotLogged(t * testing.T){
	cc := setup()
	defer cc.Close()

	cr := client{34, 100, nil}

	sr := normalpost(cr, t)

	if sr.Code != 0 {
		t.Error("Incorrect code:", sr.Code)
	}
}


//////////////////////////////////////////////////////

func setup() *connectionController{
	
	cli := clientController.ClientController()
	prob := problems.GetProblem("pruebaProblem")
	prcon := problemController.ProblemController(prob)
	
	con := ConnectionController(cli, procon)

	con.Init()

	return con
}

func encode(cr clientRequest) (io.Reader, error){
	crbytes, err := json.Marshal(cr)
	if err != nil{
		return nil, err
	}

	return bytes.NewReader(crbytes), nil
}

func decode(resp io.Reader) (*ServerResponse, error){

	dec := json.NewDecoder(resp)
	sr := &ServerResponse{}
	err := dec.Decode(sr)

	return sr, err
}

func normalPost(cr clientRequest, t* testing.T) serverResponse{
	req, err := encode(cr)
	if err != nil{
		t.Fatal(err)
	}
	
	resp, err := http.Post(url, json_type, req)
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()

	sr, err := decode(resp.Body)
	if err != nil{
		t.Fatal(err)
	}
}