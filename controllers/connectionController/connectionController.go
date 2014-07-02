package connectionController

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/nradz/DistGo/controllers/clientController"
	"github.com/nradz/DistGo/controllers/problemController"
	"github.com/nradz/DistGo/configuration"
	)

type connectionController struct{
	clientCon *clientController
	probCon *problemController
}


var(
	conf = configuration.Configuration()
	)

func ConnectionController(cli clientController.clientController, prob problemController.problemController){
	return &connectionController{cli, prob}
}

func (c *connectionController) Init(){

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

		if conf.Cors() != ""{
			w.Header().Set("Access-Control-Allow-Origin", conf.Cors())
		}

		//
		decoder := json.NewDecoder(r.Body)

		cr := ClientRequest{}

		errJSON := decoder.Decode(&cr)

		//Create the response struct. It will be modified like a pointer.
		sr := ServerResponse{}

		header := r.Header

		userAgent := header["User-Agent"]

		switch{
		
		case cr.Code == 10:
			newClient(&sr, cr, userAgent)

		case cr.Code == 20:
			newRequest(&sr, cr, userAgent)

		case cr.Code == 30:
			newResult(&sr, cr, userAgent)

		case cr.Code == 100:
			deleteClient(&sr, cr, userAgent)

		//Error cases
			
		case r.Method != "POST":
			notPostMethod(&sr, cr, userAgent)
		
		case errJSON != nil:
			notValidJSON(&sr, cr, userAgent)


		default:
			notValidCode(&sr, cr, userAgent)
		}
				
		//Response struct to string
		resByte, _ := json.Marshal(sr)
		resString := string(resByte)
		
		//Write the data in the ResponseWriter
		fmt.Fprintf(w, resString)

	})


	http.ListenAndServe(":"+conf.Port(), nil)

}

func (c *connectionController) notPostMethod(sr *ServerResponse, cr ClientRequest, userAgent []string){
	sr.Code = 0
	
	sr.Data = "The request method is not POST"
}

func (c *connectionController) notValidJSON(sr *ServerResponse, cr ClientRequest, userAgent []string){
	sr.Code = 0
	
	sr.Data = "The JSON message is not valid"
}

func (c *connectionController) notValidCode(sr *ServerResponse, cr ClientRequest, userAgent []string){
	sr.Code = 0
	
	sr.Data = "Not valid Code"
}

func (c *connectionController) newClient(sr *ServerResponse, cr ClientRequest, userAgent []string){
	
	id, err := c.clientCon.NewClient(userAgent)
	
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newClient:", err)
		return
	}

	sr.Id = id
	sr.Code = 110
	sr.Data = "Welcome!"

}

func (c *connectionController) newRequest(sr *ServerResponse, cr ClientRequest, userAgent []string){	
	sr.Id = cr.Id

	err := c.clientCon.IsLogged(cr.Id, userAgent)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest:", err)
		return
	}

	alg, update, err := c.probCon.NewRequest(cr.Id)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest:", err)
		return
	}

	sr.Data = data

	if alg != ""{
		sr.Alg = alg
		sr.code = 130
	}else{
		sr.Code = 120
	}

}

func (c *connectionController) newResult(sr *ServerResponse, cr ClientRequest, userAgent []string){

	sr.Id = cr.Id

	err := c.clientCon.IsLogged(cr.Id, userAgent)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest:", err)
		return
	}

	err := c.probCon.NewResult(cr.Id, cr.Data)
	if err != nil{
		sr.Code = 0
		fmt.println("Error en newRequest:", err)
		return
	}

	sr.Code = 140

}

func (c *connectionController) deleteClient(sr *ServerResponse, cr ClientRequest, userAgent []string){
	
	sr.Id = cr.Id

	err := c.clientCon.DeleteClient(cr.Id, userAgent)
	if err != nil{
		sr.Code = 0
		fmt.Println("Error en newRequest:", err)
		return
	}

	sr.Code = 150
	sr.Data = "Goodbye!"

	
}