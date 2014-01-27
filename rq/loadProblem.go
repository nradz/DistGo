package rq

import(
	"fmt"
	"net/http"
)

//Load a problem from database
func LoadProblem(conf map[string]string){
	if domains:= conf["CORS"]; domains != nil{
		w.Header().Set("Access-Control-Allow-Origin", domains)
	}
	
	fmt.Fprintf(w, `{'func':"windows.alert('Funciona!');"}`)

	return func(w http.ResponseWriter, r *http.Request){
		
	}
}