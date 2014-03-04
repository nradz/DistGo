package connectionController

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/nradz/DistGo/channels"
	"github.com/nradz/DistGo/configuration"
	)

var(
	clientChan = channels.ClientControlChannel()
	problemChan = channels.ProblemControlChannel()
	conf = configuration.Configuration()
	)

func ConnectionController(){

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

		switch{
		//Throwed if the request method is not POST	
		case r.Method != "POST":
			notPostMethod(&sr, cr, header)
		
		case errJSON != nil:
			notValidJSON(&sr, cr, header)

		case cr.Code == 10:
			newClient(&sr, cr, header)

		case cr.Code == 20:
			newRequest(&sr, cr, header)

		case cr.Code == 30:
			newResult(&sr, cr, header)

		case cr.Code == 100:
			deleteClient(&sr, cr, header)

		default:
			notValidCode(&sr, cr, header)
		}
				
		//Response struct to string
		resByte, _ := json.Marshal(sr)
		resString := string(resByte)
		
		//Write the data in the ResponseWriter
		fmt.Fprintf(w, resString)

	})


	http.ListenAndServe(":"+conf.Port(), nil)

}

func notPostMethod(sr *ServerResponse, cr ClientRequest, header http.Header){
	sr.Code = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "The request method is not POST"
}

func notValidJSON(sr *ServerResponse, cr ClientRequest, header http.Header){
	sr.Code = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "The JSON message is not valid"
}

func notValidCode(sr *ServerResponse, cr ClientRequest, header http.Header){
	sr.Code = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "Not valid Code"
}

func newClient(sr *ServerResponse, cr ClientRequest, header http.Header){
	
	control := &channels.ClientControlRequest{cr.Id, 10,
		header, make(chan channels.ClientControlResponse)}

	res := clientChan.Send(control)

	switch{
	
	case res.Code == 10:
		sr.Code = 110
		sr.Id = res.Id
		sr.Data = make([]string,1)
		sr.Data[0] = "welcome!"

	default:
		fmt.Println("Error en newClient: "+string(res.Code))
		sr.Id = res.Id
		sr.Code = 0
	}

}

func newRequest(sr *ServerResponse, cr ClientRequest, header http.Header){	

	control := &channels.ClientControlRequest{cr.Id, 20,
		header, make(chan channels.ClientControlResponse)}

	logged := clientChan.Send(control)

	switch{

	case logged.Code == 20:
		problemReq := &channels.ProblemControlRequest{cr.Id, 20,
			cr.Data, make(chan channels.ProblemControlResponse)}
		
		probRes := problemChan.SendRequest(problemReq)
		
		sr.Id = cr.Id
		sr.Code = probRes.Code
		sr.Alg = probRes.Alg

	//Error in the logging
	case logged.Code < 10:
		sr.Id = cr.Id
		sr.Code = 0

	//Internal Error
	default:
		fmt.Println("Error en newRequest: "+string(logged.Code))
		sr.Id = cr.Id
		sr.Code = 0
	}

}

func newResult(sr *ServerResponse, cr ClientRequest, header http.Header){
	
	control := &channels.ClientControlRequest{cr.Id, 20,
		header, make(chan channels.ClientControlResponse)}

	logged := clientChan.Send(control)

	switch{
	case logged.Code == 20:
		
		problemReq := &channels.ProblemControlRequest{cr.Id, 20,
			cr.Data, make(chan channels.ProblemControlResponse)}
		
		prob := problemChan.SendRequest(problemReq)

		sr.Id = cr.Id
		sr.Code = prob.Code
		sr.Alg = prob.Alg
		sr.Data = prob.Data

	//Error in the logging
	case logged.Code < 10:
		sr.Id = cr.Id
		sr.Code = 0

	//Internal Error
	default:
		fmt.Println("Error en newResult: "+string(logged.Code))
		sr.Id = cr.Id
		sr.Code = 0
	}

}

func deleteClient(sr *ServerResponse, cr ClientRequest, header http.Header){
	
	control := &channels.ClientControlRequest{cr.Id, 30,
		header, make(chan channels.ClientControlResponse)}

	deleted := clientChan.Send(control)

	switch{

	case deleted.Code == 30:
		sr.Code = 140
		sr.Id = cr.Id
		sr.Data = make([]string,1)
		sr.Data[0] = "Goodbye!"

	//Error in the logging
	case deleted.Code < 10:
		sr.Id = cr.Id
		sr.Code = 0

	//Internal Error
	default:
		fmt.Println("Error en newResult: "+string(deleted.Code))
		sr.Id = cr.Id
		sr.Code = 0

	}
}