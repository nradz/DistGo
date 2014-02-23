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

		case cr.Type == 10:
			newClient(&sr, cr, header)

		case cr.Type == 20:
			newRequest(&sr, cr, header)

		case cr.Type == 30:
			newResult(&sr, cr, header)

		case cr.Type == 100:
			deleteClient(&sr, cr, header)

		default:
			notValidType(&sr, cr, header)
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
	sr.Type = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "The request method is not POST"
}

func notValidJSON(sr *ServerResponse, cr ClientRequest, header http.Header){
	sr.Type = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "The JSON message is not valid"
}

func notValidType(sr *ServerResponse, cr ClientRequest, header http.Header){
	sr.Type = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "Not valid type"
}

func newClient(sr *ServerResponse, cr ClientRequest, header http.Header){
	control := &channels.ClientControlRequest{cr.Id, 10,
	 header, make(chan channels.ClientControlResponse)}	

	res := clientChan.Send(control)

	sr.Type = 110
	sr.Id = res.Id

	sr.Data = make([]string,1)
	sr.Data[0] = "welcome!"

}

func newRequest(sr *ServerResponse, cr ClientRequest, header http.Header){
	
	return
}

func newResult(sr *ServerResponse, cr ClientRequest, header http.Header){
}

func deleteClient(sr *ServerResponse, cr ClientRequest, header http.Header){
	return
}