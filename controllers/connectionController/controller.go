package connectionController

import(
	"net/http"
	"fmt"
	"encoding/json"
	)



func Controller(conf map[string]string) func(w http.ResponseWriter, r *http.Request){

	
	return func(w http.ResponseWriter, r *http.Request){

		if conf["cors"] != ""{
			w.Header().Set("Access-Control-Allow-Origin", conf["cors"])
		}

		//
		decoder := json.NewDecoder(r.Body)

		cr := ClientRequest{}

		errJSON := decoder.Decode(&cr)

		//Create the response struct. It will be modified like a pointer.
		sr := ServerResponse{}

		clientHead := r.Header

		switch{
		//Throwed if the request method is not POST	
		case r.Method != "POST":
			notPostMethod(&sr)
		
		case errJSON != nil:
			notValidJSON(&sr)

		case cr.Type == 10:
			newClient(&sr, clientHead)

		case cr.Type == 20:
			newRequest(&sr, cr, clientHead)

		case cr.Type == 30:
			newResult(&sr, cr, clientHead)

		case cr.Type == 100:
			deleteClient(&sr, cr, clientHead)

		default:
			notValidType(&sr)
		}
		
		
		//Response struct to string
		resByte, _ := json.Marshal(sr)
		resString := string(resByte)
		
		//Write the data in the ResponseWriter
		fmt.Fprintf(w, resString)

	}

}

func notPostMethod(sr *ServerResponse){
	sr.Type = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "The request method is not POST"
}

func notValidJSON(sr *ServerResponse){
	sr.Type = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "The JSON message is not valid"
}

func notValidType(sr *ServerResponse){
	sr.Type = 0
	
	sr.Data = make([]string,1)
	sr.Data[0] = "Not valid type"
}

func newClient(sr *ServerResponse, header map[string][]string){
	return
}

func newRequest(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	return
}

func newResult(sr *ServerResponse, cr ClientRequest, header map[string][]string){
}

func deleteClient(sr *ServerResponse, cr ClientRequest, header map[string][]string){
	return
}