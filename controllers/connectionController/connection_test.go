package connectionController

import(
	//"fmt"
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
	"github.com/nradz/DistGo/conf"
)

var json_type string = "application/json"

var(
	cli *clientController.ClientController
	probCon problemController.ProblemController
)

const(
	get_error = "The request method is not POST"
	json_error = "The JSON message is not valid"
	code_error = "Not valid Code"
	)

const(
	alg = `function problem(romero){
			this.mainFunc = function(data){
				romero.result(["6"]);
				romero.request();
				var up = romero.newUpdate();
				self.postMessage({'cmd':'log','message':up});
				romero.finish();
				}
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

	cr := ClientRequest{1, 0, 123, 234, nil}
	
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

	cr := ClientRequest{0, 0, 10, 0, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 110{
		t.Error("Incorrect Code: ", sr.Code)
	}

}

func TestFirstRequest(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 0, 10, 0, nil}

	//login
	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20
	cr.Key = sr.Key

	sr = normalPost(cr, t, server.URL)

	if sr.Id != cr.Id{
		t.Error("No Id:", sr.Id)
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

	if sr.Status != 1{
		t.Error("Incorrect number:", sr.Status)
	}

}


func TestNotLoggedRequest(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 0, 20, 0, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestIncorrectHeader(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 0, 10, 0, nil}

	//login
	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20
	cr.Key = sr.Key

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

	cr := ClientRequest{0, 0, 10, 0, nil}

	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 20
	cr.Key = sr.Key
	
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

	cr := ClientRequest{0, 0, 10, 0, nil}

	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Key = sr.Key
	cr.Code = 30
	
	sr = normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestNotLoggedResult(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 0, 30, 0, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0{
		t.Error("Incorrect Code:", sr.Code)
	}
}

func TestUpdate(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr1 := ClientRequest{0, 0, 10, 0, nil}
	cr2 := ClientRequest{0, 0, 10, 0, nil}

	//login
	sr1 := normalPost(cr1, t, server.URL)
	sr2 := normalPost(cr2, t, server.URL)
	
	//first request
	cr1.Id = sr1.Id
	cr2.Id = sr2.Id
	cr1.Code = 20
	cr2.Code = 20
	cr1.Key = sr1.Key
	cr2.Key = sr2.Key
	cr1.LastUpdate = sr1.Status
	cr2.LastUpdate = sr2.Status	
	
	sr1 = normalPost(cr1, t, server.URL)
	sr2 = normalPost(cr2, t, server.URL)
	cr1.LastUpdate = sr1.Status
	cr2.LastUpdate = sr2.Status	

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
	if sr2.Status != 2{
		t.Error("Incorrect number:", sr2.Status)
	}

}

func TestDelete(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	cr := ClientRequest{0, 0, 10, 0, nil}

	//login
	sr := normalPost(cr, t, server.URL)

	cr.Id = sr.Id
	cr.Code = 100
	cr.Key = sr.Key

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

	cr := ClientRequest{34, 65, 100, 0, nil}

	sr := normalPost(cr, t, server.URL)

	if sr.Code != 0 {
		t.Error("Incorrect code:", sr.Code)
	}
}

func TestMaxUserReached(t *testing.T){
	server := setup()
	defer server.Close()
	defer close()

	conf.SetMaxClients(3)

	cr1 := ClientRequest{0, 0, 10, 0, nil}
	cr2 := ClientRequest{0, 0, 10, 0, nil}
	cr3 := ClientRequest{0, 0, 10, 0, nil}

	sr1 := normalPost(cr1, t, server.URL)
	normalPost(cr2, t, server.URL)
	normalPost(cr3, t, server.URL)

	//Max limit reached
	cr4 := ClientRequest{0, 0, 10, 0, nil}
	sr4 := normalPost(cr4, t, server.URL)

	if sr4.Code != 0{
		t.Error("Incorrect code 1:", sr4.Code)
	}

	cr1.Id = sr1.Id
	cr1.Key = sr1.Key
	cr1.Code = 100

	//Delete cr1
	normalPost(cr1, t, server.URL)

	//request again
	sr4 = normalPost(cr4, t, server.URL)

	if sr4.Code == 0{
		t.Error("Incorrect code 2:", sr4.Code)
	}

}



//////////////////////////////////////////////////////

func setup() *httptest.Server{	
	cli = clientController.New()
	cli.Init()

	prob := problems.GetProblem("pruebaProblem")
	probCon = problemController.New(prob)
	if probCon == nil{
		panic("nil problemController")
	}
	probCon.Init()
	
	con := New(cli, probCon)

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