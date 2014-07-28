//Package connectionController provides a controller that manages
//the incoming requests and responses to them.
package connectionController

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/conf"
	)

//Connection Controller manages the incoming requests from the clients.
type ConnectionController struct{
	clientCon *clientController.ClientController
	probCon problemController.ProblemController
}

//New returns a new ConnectionController. It needs a clientController 
//pointer and a problemController pointer.
func New(cli *clientController.ClientController, prob problemController.ProblemController) *ConnectionController{
	return &ConnectionController{cli, prob}
}

//Init initializes the server.
func (c *ConnectionController) Init(){

	http.Handle("/", c)


	http.ListenAndServe(":"+conf.Port(), nil)

}

//ServeHTTP is the handler that will be executed by the server.
func (c *ConnectionController) ServeHTTP(w http.ResponseWriter, r *http.Request){

	if conf.Cors() != ""{
		w.Header().Set("Access-Control-Allow-Origin", conf.Cors())
	}

	//Decode json.
	decoder := json.NewDecoder(r.Body)

	cr := ClientRequest{}

	errJSON := decoder.Decode(&cr)

	//Create the response struct. It will be modified like a pointer.
	sr := ServerResponse{}

	switch{
	
	case cr.Code == 10:
		c.newClient(&sr, cr, r.Header)

	case cr.Code == 20:
		c.newRequest(&sr, cr, r.Header)

	case cr.Code == 30:
		c.newResult(&sr, cr, r.Header)

	case cr.Code == 100:
		c.deleteClient(&sr, cr, r.Header)

	//Error cases
		
	case r.Method != "POST":
		c.notPostMethod(&sr, cr, r.Header)
	
	case errJSON != nil:
		c.notValidJSON(&sr, cr, r.Header)


	default:
		c.notValidCode(&sr, cr, r.Header)
	}
			
	//Response struct to string
	resByte, _ := json.Marshal(sr)
	resString := string(resByte)
	
	//Write the data in the ResponseWriter
	fmt.Fprintf(w, resString)

}

func (c *ConnectionController) notPostMethod(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	sr.Code = 0
	
	sr.Data = "The request method is not POST"
}

func (c *ConnectionController) notValidJSON(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	sr.Code = 0
	
	sr.Data = "The JSON message is not valid"
}

func (c *ConnectionController) notValidCode(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	sr.Code = 0
	
	sr.Data = "Not valid Code"
}

func (c *ConnectionController) newClient(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	
	id, key, err := c.clientCon.NewClient(header)
	
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newClient:", err)
		return
	}
	
	sr.Id = id
	sr.Key = key
	sr.Code = 110
	sr.Data = "Welcome!"

}

func (c *ConnectionController) newRequest(sr *ServerResponse, cr ClientRequest, header map[string][]string){	
	sr.Id = cr.Id

	err := c.clientCon.IsLogged(cr.Id, cr.Key, header)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest(IsLogged):", err)
		return
	}

	alg, update, number, err := c.probCon.NewRequest(cr.Id, cr.LastUpdate)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest(NewRequest):", err)
		return
	}

	sr.Data = update
	sr.Status = number

	if alg != ""{
		sr.Alg = alg
		sr.Code = 130
	}else{
		sr.Code = 120
	}

}

func (c *ConnectionController) newResult(sr *ServerResponse, cr ClientRequest, header map[string][]string){

	sr.Id = cr.Id

	err := c.clientCon.IsLogged(cr.Id, cr.Key, header)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newResult(IsLogged):", err)
		return
	}

	err = c.probCon.NewResult(cr.Id, cr.Data, cr.LastUpdate)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newResult(NewResult):", err)
		return
	}

	sr.Code = 140

}

func (c *ConnectionController) deleteClient(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	
	sr.Id = cr.Id

	err := c.clientCon.DeleteClient(cr.Id, cr.Key, header)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest:", err)
		return
	}

	sr.Code = 150
	sr.Data = "Goodbye!"


}