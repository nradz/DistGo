package connectionController

import(
	"time"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"testing"
	"bytes"
	"io"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/problems"
)

var json_type string = "application/json"

var(
	cli *clientController.ClientController
	probCon *problemController.SimpleProblemController
)

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
	server := setup()
	defer server.Close()
	defer close()

	resp, err := http.Get(server.URL)
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()

	sr, err := decode(resp.Body)
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
	server := setup()
	defer server.Close()
	defer close()

	req := bytes.NewBufferString("Hola, niño")

	resp, err := http.Post(server.URL, "text", req)
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()

	sr, err := decode(resp.Body)
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
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{1,234,nil}
	
	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:" , sr.Code)
	}
	if sr.Data.(string) != code_error{
		t.Error("Incorrect message:", sr.Data)
	}
}

func TestNewClient(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Id == 0{
		t.Error("No id:", sr.Id)
	}

	if sr.Code != 110{
		t.Error("Incorrect Code: ", sr.Code)
	}

}

func TestFirstRequest(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	//login
	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20

	sr = normalPost(cr, t, server.URL)

	if sr.Id != cr.Id{
		t.Error("No Id:", sr.Id)
	}

	if sr.Code !=130{
		t.Error("No Code:", sr.Code)
	}

	if sr.Alg != alg{
		t.Error("No alg:", sr.Alg)
	}

	if sr.Data == nil{
		t.Error("No data:", sr.Data)
	}

}


func TestNotLoggedRequest(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 20, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestIncorrectHeader(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	//login
	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20

	json, err := encode(cr)
	if err != nil{
		t.Fatal(err)
	}

	client := &http.Client{}

	newr, err := http.NewRequest("POST", server.URL, json)
	if err != nil{
		t.Fatal(err)
	}

	newr.Header["User-Agent"] = make([]string,1)
	newr.Header["User-Agent"][0] = "Fake User-agent"
	
	resp, err := client.Do(newr)
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()

	sr, err = decode(resp.Body)
	if err != nil{
		t.Fatal(err)
	}

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}

}

func TestNewResult(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20
	
	sr = normalPost(cr, t, server.URL)

	cr.Code = 30
	cr.Data = make([]string, 1)
	cr.Data[0] = "6"
	sr = normalPost(cr, t, server.URL)

	if sr.Code != 140{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNewResultNoData(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20
	
	sr = normalPost(cr, t, server.URL)

	cr.Code = 30
	
	sr = normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNewResultNoAlg(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 30
	cr.Data = make([]string, 1)
	cr.Data[0] = "6"

	sr = normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNotLoggedResult(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 30, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestUpdate(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr1 := ClientRequest{0, 10, nil}
	cr2 := ClientRequest{0, 10, nil}

	//login
	sr1 := normalPost(cr1, t, server.URL)
	sr2 := normalPost(cr2, t, server.URL)
	
	//first request
	cr1.Id = sr1.Id
	cr2.Id = sr2.Id
	cr1.Code = 20
	cr2.Code = 20	
	
	sr1 = normalPost(cr1, t, server.URL)
	sr2 = normalPost(cr2, t, server.URL)

	//result
	cr1.Code = 30
	cr1.Data = make([]string, 1)
	cr1.Data[0] = "6"

	sr1 = normalPost(cr1, t, server.URL)

	//update
	cr2.Code = 20
	sr2 = normalPost(cr2, t, server.URL)

	if int(sr2.Data.(float64)) != 6{
		t.Error("Incorrect data:", sr2.Data)
	}

}

func TestDelete(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 10, nil}

	//login
	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 100

	sr = normalPost(cr, t, server.URL)

	if sr.Code != 150{
		t.Fatal("Incorrect Code:", sr.Code)
	}

	//Verify

	cr.Code = 20
	sr = normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestDeteleNotLogged(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{34, 100, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0 {
		t.Error("Incorrect code:", sr.Code)
	}
}


//////////////////////////////////////////////////////

func setup() *httptest.Server{	
	cli = clientController.NewClientController()
	cli.Init()

	prob := problems.GetProblem("pruebaProblem")
	probCon = problemController.NewSimpleProblemController(prob)
	probCon.Init()
	
	con := NewConnectionController(cli, probCon)

	server := httptest.NewUnstartedServer(con)	

	go server.Start()
	time.Sleep(100 *time.Millisecond)

	return server
}

func close(){
	cli.Close()
	probCon.Close()

}

func encode(cr ClientRequest) (io.Reader, error){
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

func normalPost(cr ClientRequest, t* testing.T, url string) *ServerResponse{
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

	return sr
}